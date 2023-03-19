package main

//
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/knipferrc/teacup/icons"
	"github.com/knipferrc/teacup/image"
	"github.com/knipferrc/teacup/statusbar"
)

var baseUrl string = `https://api.iwyu.com/API/baiduresou/`

type Result struct {
	Code string      `json:"code"`
	Mgs  string      `json:"mgs"`
	Data []*Trending `json:"data"`
}
type Trending struct {
	Index    int    `json:"index"`
	HotScore string `json:"hotScore"`
	Img      string `json:"img"`
	RawURL   string `json:"rawUrl"`
	Word     string `json:"word"`
	Desc     string `json:"desc"`
}

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = lipgloss.NewStyle().PaddingLeft(4).PaddingBottom(1)
	docStyle          = lipgloss.NewStyle().Margin(1, 2)
	// textStyle         = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type rowDelegate struct{}

func (d rowDelegate) Height() int {
	return 1
}

func (d rowDelegate) Spacing() int {
	return 0
}

func (d rowDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {

	return nil
}

// 渲染列表
func (d rowDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(*Trending)
	if !ok {
		return
	}
	q := (*Trending)(i)
	str := fmt.Sprintf("🔥%02d %s", q.Index, q.Word)
	if index == m.Index() {
		str = selectedItemStyle.Render("> " + str)
	} else {
		str = itemStyle.Render(str)
	}
	_, _ = fmt.Fprint(w, str)
}

type errMsg error
type TrendingList []*Trending

func (i Trending) Title() string       { return i.Word }
func (i Trending) Description() string { return i.Desc }
func (i *Trending) FilterValue() string {
	return (*Trending)(i).Word + (*Trending)(i).Desc
}

type tui struct {
	statusbar statusbar.Bubble
	image     image.Bubble
	spinner   *spinner.Model
	list      list.Model
	selected  *Trending
	loading   bool
	err       error
}

// 用bubbletea写的一个终端界面，用来选择搜索结果，然后打开链接
func newTuiModel() *tui {

	// 用自己实现的rowDelegate来渲染列表
	// c := list.New(nil, rowDelegate{}, 0, 0)
	// 用默认的rowDelegate来渲染列表
	l := list.New(nil, list.NewDefaultDelegate(), 0, 0)
	l.Title = "🔥百度热搜Trending（回车访问）"
	l.SetShowStatusBar(true)
	l.SetShowTitle(true)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	statusbarModel := statusbar.New(
		statusbar.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#ffffff"},
			Background: lipgloss.AdaptiveColor{Dark: "#cc241d", Light: "#cc241d"},
		},
		statusbar.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#ffffff"},
			Background: lipgloss.AdaptiveColor{Dark: "#3c3836", Light: "#3c3836"},
		},
		statusbar.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#ffffff"},
			Background: lipgloss.AdaptiveColor{Dark: "#A550DF", Light: "#A550DF"},
		},
		statusbar.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#ffffff"},
			Background: lipgloss.AdaptiveColor{Dark: "#6124DF", Light: "#6124DF"},
		},
	)
	statusbarModel.SetContent(
		"测试",
		"我是进度条",
		"123",
		"测试",
	)
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	// TODO Implement a progressive loading list
	return &tui{
		image:     image.New(true, true, lipgloss.AdaptiveColor{Light: "#000000", Dark: "#ffffff"}),
		statusbar: statusbarModel,
		spinner:   &s,
		list:      l,
		loading:   true,
	}
}
func (m *tui) Selected() *Trending {
	return m.selected
}

func (m *tui) Init() tea.Cmd {

	loading := func() tea.Msg {
		return spinner.TickMsg{}
	}
	init := func() tea.Msg {
		// 用http从baseUrl获取搜索结果映射到result
		resp, err := http.Get(baseUrl)
		defer resp.Body.Close()
		// 得到请求的status code
		if err != nil {
			return err
		}

		var res Result
		if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return err
		}
		items := make([]list.Item, len(res.Data))
		for i, q := range res.Data {
			items[i] = q
		}
		m.list.SetItems(items)
		m.loading = false
		return nil
	}
	// 注册加载中的动画和初始化事件
	return tea.Batch(loading, init, m.image.SetFileName("logo1.png"))
}

// 下载远程图片到当前目录下cache文件夹
func (m *tui) downloadImg(url string) tea.Cmd {
	// 获取图片文件名
	filename := filepath.Base(url)
	// 拼接本地缓存路径
	cachePath := filepath.Join("cache", filename)
	if _, err := os.Stat("cache"); os.IsNotExist(err) {
		err = os.Mkdir("cache", 0755)
		if err != nil {
			return func() tea.Msg { return err }
		}
	}
	// 如果本地已经存在该文件，直接返回文件路径
	// if _, err := os.Stat(cachePath); err == nil {
	// 	return m.image.SetFileName(cachePath)
	// }

	// 否则，下载图片并保存到本地缓存
	resp, err := http.Get(url)
	if err != nil {
		return func() tea.Msg { return err }
	}
	defer resp.Body.Close()

	out, err := os.Create(cachePath)
	if err != nil {
		return func() tea.Msg { return err }
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return func() tea.Msg { return err }
	}

	return m.image.SetFileName(cachePath)
}

func (m *tui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.list.SelectedItem() != nil {
				m.selected = (*Trending)(m.list.SelectedItem().(*Trending))
				return m, tea.Quit
			}
		case " ":
			// cmd := m.image.SetFileName(fmt.Sprintf("./logo%d.png", m.list.Index()%2))
			// m.image.SetBorderColor(lipgloss.AdaptiveColor{Dark: "#F25D94", Light: "#F25D94"})
			return m, m.downloadImg(m.list.SelectedItem().(*Trending).Img)
		}
	case tea.WindowSizeMsg:
		m.statusbar.SetSize(msg.Width)
		m.list.SetSize(msg.Width/2, msg.Height-m.statusbar.Height)
		resizeImgCmd := m.image.SetSize(msg.Width-m.list.Width(), msg.Height-m.statusbar.Height)
		return m, tea.Batch(resizeImgCmd)
	case spinner.TickMsg:
		// 用来显示加载中的动画
		if m.loading == false {
			return m, nil
		}
		spinner, cmd := m.spinner.Update(msg)
		m.spinner = &spinner
		return m, cmd

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.list, _ = m.list.Update(msg)
	m.image, _ = m.image.Update(msg)

	if m.loading {
		return m, nil
	}

	logoText := fmt.Sprintf("%s %s", icons.IconDef["dir"].GetGlyph(), "百度热搜")
	m.statusbar.SetContent(
		m.list.SelectedItem().(*Trending).Word,
		m.list.SelectedItem().(*Trending).Desc,
		fmt.Sprintf("%d/%d", m.list.Index(), len(m.list.Items())),
		logoText,
	)
	return m, nil
}

func (m *tui) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	if m.loading {
		return fmt.Sprintf("\n\n   %s 正在加载数据\n %s \n", m.spinner.View(), m.statusbar.View())
	} else {
		return "\n" +
			lipgloss.JoinVertical(lipgloss.Top,
				lipgloss.JoinHorizontal(lipgloss.Left, m.list.View(), m.image.View()),
				m.statusbar.View(),
			)
	}
}

// 写个go写个终端打开链接的函数
func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	case "linux":
		cmd = "xdg-open"
	default:
		return fmt.Errorf("unsupported platform")
	}

	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func main() {
	m := newTuiModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(0)
	}
	if m.Selected() == nil {
		os.Exit(0)
	}
	q := m.Selected()
	//用浏览器打开q的链接
	open(string(q.RawURL))
}
