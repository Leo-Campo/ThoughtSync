/*
Copyright © 2024 Leonardo Campitelli leonardo932.campitelli@gmail.com
*/
package cmd

import (
	"ThoughtSync/cmd/config"
	"ThoughtSync/cmd/date"
	"ThoughtSync/cmd/editor"
	"ThoughtSync/cmd/path"
	"fmt"
	"time"

	gopath "path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// OpenTodayNote opens the today note with name filename in the vault directory
// vaultJournalPath using the editor provider
func OpenTodayNote(editor editor.Editor, vaultJournalPath, filename string) error {
	filenameMd := fmt.Sprintf("%s.md", filename)
	err := path.EnsurePresent(vaultJournalPath, filenameMd)
	if err != nil {
		return fmt.Errorf("failed to ensure present: %w", err)
	}
	filePath := gopath.Join(vaultJournalPath, filenameMd)
	err = editor.Edit(filePath)
	if err != nil {
		return fmt.Errorf("error in editing file: %w", err)
	}
	return nil
}

func init() {
	editor := editor.NewEditor()
	todayCmd := &cobra.Command{
		Use:   "today",
		Short: "Quickly edit the journal note for today",
		RunE: func(cmd *cobra.Command, args []string) error {
			vaultPath := viper.GetString(config.VAULT_KEY)
			format := viper.GetString(config.JOURNAL_NOTE_FORMAT_KEY)
			journalDir := viper.GetString(config.JOURNAL_DIRECTORY_KEY)
			vaultJournalPath := gopath.Join(vaultPath, journalDir)
			filename, err := date.Format(time.Now(), format)
			if err != nil {
				return fmt.Errorf("error getting journal filename: %w", err)
			}
			return OpenTodayNote(editor, vaultJournalPath, filename)
		},
	}
	RootCmd.AddCommand(todayCmd)
}
