package cmd

import (
    "bufio"
    "embed"
    "errors"
    "fmt"
    "os"
    "strconv"
    "strings"
    "text/template"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"

    log "github.com/cstaaben/go-log"

    "github.com/cstaaben/adventofcode/internal/convert"
)

var (
    //go:embed templates
    templates embed.FS

    funcMap = map[string]any{
        "NumToWord":     convert.NumberToWord,
        "KebabToLower":  convert.KebabToLower,
        "KebabToPascal": convert.KebabToPascal,
        "KebabToTitle":  convert.KebabToTitle,
        "CamelToPascal": convert.CamelToPascal,
    }
)

func newAddCommand() *cobra.Command {
    addCmd := &cobra.Command{
        Use:           "add_day",
        Short:         "add a new stubbed day subcommand for Advent of Code",
        RunE:          addDay,
        Args:          cobra.MaximumNArgs(1),
        SilenceErrors: true,
    }

    addCmd.Flags().Bool("dry-run", false,
        "Enable dry-run mode to run through all generations without creating any files.")

    return addCmd
}

func addDay(_ *cobra.Command, args []string) error {
    var (
        day string

        templateArgs = make(map[string]any)
        logger       = log.New(os.Stderr)
    )

    if viper.GetBool("debug") {
        logger.WithDebug()
    }

    // validate argument is an integer before assigning to template arguments
    if len(args) == 1 {
        n, err := strconv.Atoi(args[0])
        if err != nil {
            return fmt.Errorf("validating argument: %w", err)
        }

        templateArgs["addDay"] = convert.NumberToWord(n)
    }

    t, err := template.New("add").Funcs(funcMap).ParseFS(templates, "templates/*.gotmpl")
    if err != nil {
        logger.Error("unable to parse template files:", err)
        return err
    }

    days, err := determineExistingDays()
    if err != nil {
        return fmt.Errorf("determining existing day commands: %w", err)
    }

    if day == "" {
        lastDay := len(days) // determineExistingDays doesn't sort days
        nextDay := convert.NumberToWord(lastDay + 1)

        days = append(days, nextDay)
        templateArgs["addDay"] = nextDay
    } else {
        days = []string{day}
    }

    templateArgs["days"] = days

    // logger.Debugf("found %d existing days; using %v as next day", len(days), templateArgs["addDay"])
    logger.Debugf("days: %v\n", days)

    err = executeTemplateToFile(t.Lookup("days_root"), "cmd/days.go", templateArgs)
    if err != nil {
        return fmt.Errorf("creating days.go: %w", err)
    }

    logger.Debug("re-generated days.go to add new subcommand")

    filename := fmt.Sprintf("internal/days/%s.go", convert.KebabToCamel(templateArgs["addDay"].(string)))
    err = executeTemplateToFile(t.Lookup("subcommand"), filename, templateArgs)
    if err != nil {
        return fmt.Errorf("adding subcommand: %w", err)
    }

    err = os.Mkdir(fmt.Sprintf("input/%s", convert.KebabToLower(templateArgs["addDay"].(string))), 0755)
    if err != nil {
        return fmt.Errorf("creating input subdirectory: %w", err)
    }

    logger.Debugf("generated internal/days/%s\n", filename)

    return nil
}

func determineExistingDays() ([]string, error) {
    info, err := os.ReadDir("internal/days")
    if err != nil {
        return nil, fmt.Errorf("reading internal/days: %w", err)
    }

    days := make([]string, 0)
    for _, entry := range info {
        days = append(days, strings.TrimSuffix(entry.Name(), ".go"))
    }

    return days, nil
}

func executeTemplateToFile(t *template.Template, filename string, args map[string]interface{}) error {
    if t == nil {
        return errors.New("unable to find template")
    }

    var writer *bufio.Writer
    if !viper.GetBool("dry-run") {
        file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
        if err != nil {
            return fmt.Errorf("opening file: %w", err)
        }
        defer file.Close() // nolint:errcheck

        writer = bufio.NewWriter(file)
    } else {
        writer = bufio.NewWriter(os.Stdout)
    }

    err := t.Execute(writer, args)
    if err != nil {
        return fmt.Errorf("executing template: %w", err)
    }
    if err = writer.Flush(); err != nil {
        return fmt.Errorf("flushing file contents: %w", err)
    }

    return nil
}
