package server

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/ya-breeze/geekbudgetbe/pkg/config"
	"github.com/ya-breeze/geekbudgetbe/pkg/database"
	"github.com/ya-breeze/geekbudgetbe/pkg/generated/goserver"
)

func Server(logger *slog.Logger, cfg *config.Config) error {
	storage := database.NewStorage(logger, cfg)
	if err := storage.Open(); err != nil {
		return fmt.Errorf("Failed to open storage: %w", err)
	}

	logger.Info("Starting GeekBudget server...")

	if cfg.Users != "" {
		logger.Info("Creating users...")
		users := strings.Split(cfg.Users, ",")
		for _, user := range users {
			tokens := strings.Split(user, ":")
			if len(tokens) != 2 {
				return fmt.Errorf("Invalid user format: %s", user)
			}

			user, err := storage.GetUser(tokens[0])
			if err != nil {
				return fmt.Errorf("Failed to reading user from DB: %w", err)
			}
			if user != nil {
				logger.Info(fmt.Sprintf("Updating password for user %q", tokens[0]))
				user.HashedPassword = tokens[1]
				if err := storage.PutUser(user); err != nil {
					return fmt.Errorf("Failed to update user: %w", err)
				}
			} else {
				logger.Info(fmt.Sprintf("Creating user %q", tokens[0]))
				if err := storage.CreateUser(tokens[0], tokens[1]); err != nil {
					return fmt.Errorf("Failed to create user: %w", err)
				}
			}
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, finishCham, err := goserver.Serve(ctx, logger, cfg)
	if err != nil {
		return fmt.Errorf("Failed to serve: %w", err)
	}

	// Wait for an interrupt signal
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-stopChan
	logger.Info("Received signal. Shutting down server...")

	// Stop the server
	cancel()
	<-finishCham
	return nil

	// userId := "123e4567-e89b-12d3-a456-426614174000"

	// acc := &goserver.AccountNoId{
	// 	Name:        "Test Account",
	// 	Type:        "CHECKING",
	// 	Description: "Test Account Description",
	// }

	// account, err := storage.CreateAccount(userId, acc)
	// if err != nil {
	// 	return fmt.Errorf("Failed to create account: %w", err)
	// }

	// logger.With("account", account).Info("Account created")

	// accounts, err := storage.GetAccounts(userId)
	// if err != nil {
	// 	return fmt.Errorf("Failed to get accounts: %w", err)
	// }

	// logger.With("accounts", accounts).Info(fmt.Sprintf("Accounts retrieved: %d", len(accounts)))

	// logger.Info("GeekBudget stopped")

	// return nil
}
