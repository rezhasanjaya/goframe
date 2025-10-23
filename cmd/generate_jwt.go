package cmd

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var generateJWTCmd = &cobra.Command{
	Use:   "generate:jwt",
	Short: "Generate a new JWT secret and update .env",
	Run: func(cmd *cobra.Command, args []string) {
		secret, err := generateRandomSecret(32)
		if err != nil {
			fmt.Println("❌ Failed to generate secret:", err)
			return
		}

		envPath := ".env"
		content, err := os.ReadFile(envPath)
		if err != nil {
			fmt.Println("❌ Failed to read .env:", err)
			return
		}

		lines := strings.Split(string(content), "\n")
		updated := false
		for i, line := range lines {
			if strings.HasPrefix(line, "JWT_SECRET=") {
				lines[i] = "JWT_SECRET=" + secret
				updated = true
			} else if strings.HasPrefix(line, "ACCESS_TOKEN_TTL_MIN=") {
				lines[i] = "ACCESS_TOKEN_TTL_MIN=15"
				updated = true
			} else if strings.HasPrefix(line, "REFRESH_TOKEN_TTL_H=") {
				lines[i] = "REFRESH_TOKEN_TTL_H=24"
				updated = true
			}
		}

		if !updated {
			lines = append(lines,
				"JWT_SECRET="+secret,
				"ACCESS_TOKEN_TTL_MIN=15",
				"REFRESH_TOKEN_TTL_H=24",
			)
		}

		if err := os.WriteFile(envPath, []byte(strings.Join(lines, "\n")), 0644); err != nil {
			fmt.Println("❌ Failed to write .env:", err)
			return
		}

		fmt.Println("✅ JWT secret and TTL updated in .env")
		fmt.Println("JWT_SECRET:", secret)
	},
}

func generateRandomSecret(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func init() {
	rootCmd.AddCommand(generateJWTCmd)
}
