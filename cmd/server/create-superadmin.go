package main

import (
	"fmt"
	"github.com/ImSingee/1man-verify/model"
	"github.com/ImSingee/1man-verify/service/user"
	"github.com/spf13/cobra"
)

var saUsername string
var saPassword string
var saEmail string

var cmdCreateSuperAdmin = &cobra.Command{
	Use:   "create-superadmin",
	Short: "Create a superadmin",
	Long:  `Create a superadmin`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if saPassword == "" {
			return fmt.Errorf("password is required")
		}

		_, err := setupDB()
		if err != nil {
			return err
		}

		u := &model.User{
			Username:     saUsername,
			RealPassword: saPassword,
			Email:        saEmail,
			Role:         model.RoleSuperAdmin,
		}

		err = user.CreateUser(u)
		if err != nil {
			return err
		}

		fmt.Printf("Superadmin %s created\n", saUsername)
		fmt.Printf("ID = %d\n", u.ID)

		return nil
	},
}

func init() {
	app.AddCommand(cmdCreateSuperAdmin)

	cmdCreateSuperAdmin.Flags().StringVarP(&saUsername, "username", "u", "admin", "")
	cmdCreateSuperAdmin.Flags().StringVarP(&saEmail, "email", "e", "admin@3600.academy", "")
	cmdCreateSuperAdmin.Flags().StringVarP(&saPassword, "password", "p", "", "")
}
