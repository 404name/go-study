package cmd

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/404name/go-study/demos/cobra/public"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// VersionCmd represents the version command
var ServerCmd = &cobra.Command{
	Use:   "service",
	Short: "启动web服务",
	Run: func(cmd *cobra.Command, args []string) {
		// 创建一个默认的 Gin 引擎
		checkStaticFiles()
		r := gin.Default()

		// 添加一个路由处理函数

		r.GET("/hello", func(c *gin.Context) {
			c.String(200, "Hello, World!"+fmt.Sprintf(`
			Version: %s
			staticDir %s
			flag: %v
			`,
				version, staticDir, flag))
		})

		// 静态文件服务
		r.Static("/static", staticDir)
		// 其他静态文件处理
		r.GET("/:filename", func(c *gin.Context) {
			filename := c.Param("filename")
			c.File(staticDir + "/" + filename)
		})

		// 启动引擎，监听 8080 端口
		r.Run(":8080")
	},
}

func checkStaticFiles() {
	// 如果静态文件夹不存在，则从嵌入式文件系统中提取文件
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		fmt.Printf("提取静态文件到 %s...\n", staticDir)
		err := extractStaticFiles(public.StaticContent, staticDir)
		if err != nil {
			fmt.Printf("提取静态文件失败：%s\n", err)
			os.Exit(1)
		}
	}
}

// 从嵌入式文件系统中提取静态文件
func extractStaticFiles(content embed.FS, dest string) error {
	err := fs.WalkDir(content, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			data, err := content.ReadFile(path)
			if err != nil {
				return err
			}
			// 将第一个参数设为 dest 的父目录，这样就可以避免 dist/dist 的问题
			err = os.MkdirAll(filepath.Dir(filepath.Join(dest, "..", path)), 0755)
			if err != nil {
				return err
			}
			err = os.WriteFile(filepath.Join(dest, "..", path), data, 0644)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
