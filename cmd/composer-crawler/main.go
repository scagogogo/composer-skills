package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/scagogogo/composer-crawler/pkg/client"
	"github.com/scagogogo/composer-crawler/pkg/domain"
	"github.com/spf13/cobra"
)

var (
	// 全局选项
	timeout    int
	baseURL    string
	repoURL    string
	username   string
	apiToken   string
	outputJSON bool

	// 用于显示的彩色输出
	errorColor   = color.New(color.FgRed, color.Bold)
	successColor = color.New(color.FgGreen, color.Bold)
	infoColor    = color.New(color.FgCyan)
	warnColor    = color.New(color.FgYellow)
	titleColor   = color.New(color.FgMagenta, color.Bold)
)

// 创建ComposerClient实例
func createClient() *client.ComposerClient {
	options := []client.ComposerClientOption{}

	if baseURL != "" {
		options = append(options, client.WithBaseURL(baseURL))
	}
	if repoURL != "" {
		options = append(options, client.WithRepoURL(repoURL))
	}
	if username != "" && apiToken != "" {
		options = append(options, client.WithAPICredentials(username, apiToken))
	}

	return client.NewComposerClient(time.Duration(timeout)*time.Second, options...)
}

// 输出结果，支持JSON或者友好的文本格式
func printResult(result interface{}) {
	if outputJSON {
		// 如果需要JSON输出，则使用内建的JSON编码
		if err := json.NewEncoder(os.Stdout).Encode(result); err != nil {
			errorColor.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// 根据结果类型进行友好的文本格式化输出
	// 具体实现将在各个命令中处理
	fmt.Printf("%+v\n", result)
}

func main() {
	// 创建根命令
	rootCmd := &cobra.Command{
		Use:   "composer-crawler",
		Short: "A CLI tool for interacting with Composer/Packagist API",
		Long: `Composer Crawler is a command-line tool for interacting with the Composer/Packagist API.
It allows you to search packages, get package information, statistics, and security advisories.`,
	}

	// 添加全局选项
	rootCmd.PersistentFlags().IntVar(&timeout, "timeout", 60, "HTTP request timeout in seconds")
	rootCmd.PersistentFlags().StringVar(&baseURL, "base-url", "https://packagist.org", "Base URL for Packagist API")
	rootCmd.PersistentFlags().StringVar(&repoURL, "repo-url", "https://repo.packagist.org", "Repository URL for Composer API")
	rootCmd.PersistentFlags().StringVar(&username, "username", "", "Username for Packagist API")
	rootCmd.PersistentFlags().StringVar(&apiToken, "api-token", "", "API token for Packagist API")
	rootCmd.PersistentFlags().BoolVar(&outputJSON, "json", false, "Output results as JSON")

	// 添加子命令
	rootCmd.AddCommand(createPackageCmd())
	rootCmd.AddCommand(createRepoCmd())
	rootCmd.AddCommand(createSearchCmd())
	rootCmd.AddCommand(createSecurityCmd())
	rootCmd.AddCommand(createChangesCmd())
	rootCmd.AddCommand(createManageCmd())

	// 执行命令
	if err := rootCmd.Execute(); err != nil {
		errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// createPackageCmd 创建package相关命令
func createPackageCmd() *cobra.Command {
	packageCmd := &cobra.Command{
		Use:   "package",
		Short: "Package related commands",
		Long:  `Commands for getting information about packages.`,
	}

	// 子命令: package info
	infoCmd := &cobra.Command{
		Use:   "info [package]",
		Short: "Get package information",
		Long:  `Get detailed information about a package from Packagist.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			packageName := args[0]
			c := createClient()

			infoColor.Fprintf(os.Stderr, "Fetching package information for %s...\n", packageName)
			packageInfo, err := c.GetPackage(packageName)
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			successColor.Fprintln(os.Stderr, "Successfully fetched package information!")
			printResult(packageInfo)
		},
	}

	// 子命令: package stats
	statsCmd := &cobra.Command{
		Use:   "stats [package]",
		Short: "Get package statistics",
		Long:  `Get download statistics for a package from Packagist.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			packageName := args[0]
			c := createClient()

			infoColor.Fprintf(os.Stderr, "Fetching statistics for %s...\n", packageName)
			stats, err := c.GetPackageStats(packageName)
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			successColor.Fprintln(os.Stderr, "Successfully fetched package statistics!")
			printResult(stats)
		},
	}

	// 子命令: package v2-metadata
	v2MetadataCmd := &cobra.Command{
		Use:   "v2-metadata [package]",
		Short: "Get package with V2 metadata",
		Long:  `Get package information with Composer V2 metadata format.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			packageName := args[0]
			c := createClient()

			infoColor.Fprintf(os.Stderr, "Fetching V2 metadata for %s...\n", packageName)
			data, err := c.GetPackageWithV2Metadata(packageName)
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			successColor.Fprintln(os.Stderr, "Successfully fetched V2 metadata!")
			fmt.Println(string(data))
		},
	}

	// 子命令: package dev-versions
	devVersionsCmd := &cobra.Command{
		Use:   "dev-versions [package]",
		Short: "Get package development versions",
		Long:  `Get development versions of a package.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			packageName := args[0]
			c := createClient()

			infoColor.Fprintf(os.Stderr, "Fetching development versions for %s...\n", packageName)
			data, err := c.GetPackageDevVersions(packageName)
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			successColor.Fprintln(os.Stderr, "Successfully fetched development versions!")
			fmt.Println(string(data))
		},
	}

	packageCmd.AddCommand(infoCmd, statsCmd, v2MetadataCmd, devVersionsCmd)
	return packageCmd
}

// createRepoCmd 创建repo相关命令
func createRepoCmd() *cobra.Command {
	repoCmd := &cobra.Command{
		Use:   "repo",
		Short: "Repository related commands",
		Long:  `Commands for getting information about the Packagist repository.`,
	}

	// 子命令: repo stats
	statsCmd := &cobra.Command{
		Use:   "stats",
		Short: "Get repository statistics",
		Long:  `Get statistics about the Packagist repository.`,
		Run: func(cmd *cobra.Command, args []string) {
			c := createClient()

			infoColor.Fprintln(os.Stderr, "Fetching repository statistics...")
			stats, err := c.GetStatistics()
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			successColor.Fprintln(os.Stderr, "Successfully fetched repository statistics!")
			printResult(stats)
		},
	}

	// 子命令: repo list
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all packages",
		Long:  `List all packages in the Packagist repository.`,
		Run: func(cmd *cobra.Command, args []string) {
			c := createClient()

			infoColor.Fprintln(os.Stderr, "Fetching package list...")
			list, err := c.ListPackages()
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			successColor.Fprintf(os.Stderr, "Successfully fetched %d packages!\n", len(list.PackageNames))
			printResult(list)
		},
	}

	// 子命令: repo list-vendor
	var vendor string
	listVendorCmd := &cobra.Command{
		Use:   "list-vendor",
		Short: "List packages by vendor",
		Long:  `List packages from a specific vendor in the Packagist repository.`,
		Run: func(cmd *cobra.Command, args []string) {
			if vendor == "" {
				errorColor.Fprintln(os.Stderr, "Error: vendor is required")
				os.Exit(1)
			}

			c := createClient()

			infoColor.Fprintf(os.Stderr, "Fetching packages for vendor %s...\n", vendor)
			list, err := c.ListPackagesByVendor(vendor)
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			successColor.Fprintf(os.Stderr, "Successfully fetched %d packages!\n", len(list.PackageNames))
			printResult(list)
		},
	}
	listVendorCmd.Flags().StringVar(&vendor, "vendor", "", "Vendor name (required)")
	listVendorCmd.MarkFlagRequired("vendor")

	// 子命令: repo list-type
	var packageType string
	listTypeCmd := &cobra.Command{
		Use:   "list-type",
		Short: "List packages by type",
		Long:  `List packages of a specific type in the Packagist repository.`,
		Run: func(cmd *cobra.Command, args []string) {
			if packageType == "" {
				errorColor.Fprintln(os.Stderr, "Error: type is required")
				os.Exit(1)
			}

			c := createClient()

			infoColor.Fprintf(os.Stderr, "Fetching packages of type %s...\n", packageType)
			list, err := c.ListPackagesByType(packageType)
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			successColor.Fprintf(os.Stderr, "Successfully fetched %d packages!\n", len(list.PackageNames))
			printResult(list)
		},
	}
	listTypeCmd.Flags().StringVar(&packageType, "type", "", "Package type (required)")
	listTypeCmd.MarkFlagRequired("type")

	// 子命令: repo list-with-data
	var fields []string
	listWithDataCmd := &cobra.Command{
		Use:   "list-with-data",
		Short: "List packages with additional data",
		Long:  `List packages with additional data fields.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(fields) == 0 {
				errorColor.Fprintln(os.Stderr, "Error: at least one field is required")
				os.Exit(1)
			}

			c := createClient()

			infoColor.Fprintf(os.Stderr, "Fetching packages with additional data (%v)...\n", fields)
			list, err := c.ListPackagesWithData(fields)
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			successColor.Fprintf(os.Stderr, "Successfully fetched packages with data!\n")
			printResult(list)
		},
	}
	listWithDataCmd.Flags().StringSliceVar(&fields, "fields", []string{}, "Data fields to include (e.g. repository,type)")
	listWithDataCmd.MarkFlagRequired("fields")

	// 子命令: repo popular
	var perPage int
	popularCmd := &cobra.Command{
		Use:   "popular",
		Short: "List popular packages",
		Long:  `List popular packages from the Packagist repository.`,
		Run: func(cmd *cobra.Command, args []string) {
			c := createClient()

			infoColor.Fprintf(os.Stderr, "Fetching popular packages (limit: %d)...\n", perPage)
			list, err := c.ListPopularPackages(perPage)
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			successColor.Fprintf(os.Stderr, "Successfully fetched %d popular packages!\n", len(list.Packages))
			printResult(list)
		},
	}
	popularCmd.Flags().IntVar(&perPage, "per-page", 100, "Number of results per page")

	repoCmd.AddCommand(statsCmd, listCmd, listVendorCmd, listTypeCmd, listWithDataCmd, popularCmd)
	return repoCmd
}

// createSearchCmd 创建search相关命令
func createSearchCmd() *cobra.Command {
	searchCmd := &cobra.Command{
		Use:   "search",
		Short: "Search for packages",
		Long:  `Commands for searching packages in the Packagist repository.`,
	}

	var perPage, page int

	// 子命令: search query
	queryCmd := &cobra.Command{
		Use:   "query [query]",
		Short: "Search packages by query",
		Long:  `Search for packages using a query string.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			query := args[0]
			c := createClient()

			infoColor.Fprintf(os.Stderr, "Searching for packages matching '%s'...\n", query)
			results, err := c.SearchPackages(query, perPage, page)
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			successColor.Fprintf(os.Stderr, "Found %d packages (total: %d)!\n", len(results.Results), results.Total)
			printResult(results)
		},
	}
	queryCmd.Flags().IntVar(&perPage, "per-page", 15, "Number of results per page")
	queryCmd.Flags().IntVar(&page, "page", 1, "Page number")

	// 子命令: search tags
	var tags []string
	tagsCmd := &cobra.Command{
		Use:   "tags",
		Short: "Search packages by tags",
		Long:  `Search for packages by tags.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(tags) == 0 {
				errorColor.Fprintln(os.Stderr, "Error: at least one tag is required")
				os.Exit(1)
			}

			c := createClient()

			infoColor.Fprintf(os.Stderr, "Searching for packages with tags %v...\n", tags)
			results, err := c.SearchPackagesByTags(tags, perPage, page)
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			successColor.Fprintf(os.Stderr, "Found %d packages (total: %d)!\n", len(results.Results), results.Total)
			printResult(results)
		},
	}
	tagsCmd.Flags().StringSliceVar(&tags, "tags", []string{}, "Tags to search for")
	tagsCmd.Flags().IntVar(&perPage, "per-page", 15, "Number of results per page")
	tagsCmd.Flags().IntVar(&page, "page", 1, "Page number")
	tagsCmd.MarkFlagRequired("tags")

	// 子命令: search type
	var packageType string
	typeCmd := &cobra.Command{
		Use:   "type [query]",
		Short: "Search packages by type",
		Long:  `Search for packages by type.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			query := args[0]

			if packageType == "" {
				errorColor.Fprintln(os.Stderr, "Error: type is required")
				os.Exit(1)
			}

			c := createClient()

			infoColor.Fprintf(os.Stderr, "Searching for packages matching '%s' with type '%s'...\n", query, packageType)
			results, err := c.SearchPackagesByType(query, packageType, perPage, page)
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			successColor.Fprintf(os.Stderr, "Found %d packages (total: %d)!\n", len(results.Results), results.Total)
			printResult(results)
		},
	}
	typeCmd.Flags().StringVar(&packageType, "type", "", "Package type (required)")
	typeCmd.Flags().IntVar(&perPage, "per-page", 15, "Number of results per page")
	typeCmd.Flags().IntVar(&page, "page", 1, "Page number")
	typeCmd.MarkFlagRequired("type")

	searchCmd.AddCommand(queryCmd, tagsCmd, typeCmd)
	return searchCmd
}

// createSecurityCmd 创建security相关命令
func createSecurityCmd() *cobra.Command {
	securityCmd := &cobra.Command{
		Use:   "security",
		Short: "Security related commands",
		Long:  `Commands for getting security information about packages.`,
	}

	// 子命令: security advisories
	advisoriesCmd := &cobra.Command{
		Use:   "advisories",
		Short: "Get security advisories",
		Long:  `Get security advisories for all packages.`,
		Run: func(cmd *cobra.Command, args []string) {
			c := createClient()

			infoColor.Fprintln(os.Stderr, "Fetching security advisories...")
			advisories, err := c.GetSecurityAdvisories()
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			successColor.Fprintln(os.Stderr, "Successfully fetched security advisories!")
			printResult(advisories)
		},
	}

	// 子命令: security package
	var packages []string
	packageCmd := &cobra.Command{
		Use:   "package",
		Short: "Get security advisories for packages",
		Long:  `Get security advisories for specific packages.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(packages) == 0 {
				errorColor.Fprintln(os.Stderr, "Error: at least one package is required")
				os.Exit(1)
			}

			c := createClient()

			infoColor.Fprintf(os.Stderr, "Fetching security advisories for packages %v...\n", packages)
			advisories, err := c.GetSecurityAdvisoriesForPackages(packages)
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			successColor.Fprintln(os.Stderr, "Successfully fetched security advisories!")
			printResult(advisories)
		},
	}
	packageCmd.Flags().StringSliceVar(&packages, "packages", []string{}, "Packages to get advisories for")
	packageCmd.MarkFlagRequired("packages")

	// 子命令: security since
	var sinceDays int
	sinceCmd := &cobra.Command{
		Use:   "since",
		Short: "Get security advisories since a timestamp",
		Long:  `Get security advisories updated since a specific timestamp.`,
		Run: func(cmd *cobra.Command, args []string) {
			c := createClient()

			updatedSince := time.Now().AddDate(0, 0, -sinceDays)
			infoColor.Fprintf(os.Stderr, "Fetching security advisories updated since %s...\n", updatedSince.Format("2006-01-02"))
			advisories, err := c.GetSecurityAdvisoriesSince(updatedSince)
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			successColor.Fprintln(os.Stderr, "Successfully fetched security advisories!")
			printResult(advisories)
		},
	}
	sinceCmd.Flags().IntVar(&sinceDays, "days", 30, "Number of days to look back")

	securityCmd.AddCommand(advisoriesCmd, packageCmd, sinceCmd)
	return securityCmd
}

// createChangesCmd 创建changes相关命令
func createChangesCmd() *cobra.Command {
	var since int64

	changesCmd := &cobra.Command{
		Use:   "changes",
		Short: "Get package changes",
		Long:  `Get information about package changes.`,
		Run: func(cmd *cobra.Command, args []string) {
			c := createClient()

			infoColor.Fprintf(os.Stderr, "Fetching package changes since %d...\n", since)
			changes, err := c.GetPackageChanges(cmd.Context(), since)
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			if changes.Error != "" {
				warnColor.Fprintf(os.Stderr, "API Warning: %s\n", changes.Error)
			}

			successColor.Fprintln(os.Stderr, "Successfully fetched package changes!")
			printResult(changes)
		},
	}

	changesCmd.Flags().Int64Var(&since, "since", 0, "Timestamp to get changes since")

	return changesCmd
}

// createManageCmd 创建manage相关命令
func createManageCmd() *cobra.Command {
	manageCmd := &cobra.Command{
		Use:   "manage",
		Short: "Package management commands",
		Long:  `Commands for managing packages (requires API credentials).`,
	}

	// 子命令: manage create
	var repository string
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new package",
		Long:  `Create a new package in Packagist (requires API credentials).`,
		Run: func(cmd *cobra.Command, args []string) {
			if repository == "" {
				errorColor.Fprintln(os.Stderr, "Error: repository URL is required")
				os.Exit(1)
			}

			c := createClient()

			// 检查是否设置了凭据
			if username == "" || apiToken == "" {
				errorColor.Fprintln(os.Stderr, "Error: username and API token are required")
				os.Exit(1)
			}

			infoColor.Fprintf(os.Stderr, "Creating package from repository %s...\n", repository)
			result, err := c.CreatePackage(cmd.Context(), &domain.PackageCreateRequest{
				Repository: repository,
			})
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			successColor.Fprintln(os.Stderr, "Successfully created package!")
			printResult(result)
		},
	}
	createCmd.Flags().StringVar(&repository, "repository", "", "Repository URL (required)")
	createCmd.MarkFlagRequired("repository")

	// 子命令: manage edit
	var packageName string
	editCmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit an existing package",
		Long:  `Edit an existing package in Packagist (requires API credentials).`,
		Run: func(cmd *cobra.Command, args []string) {
			if packageName == "" {
				errorColor.Fprintln(os.Stderr, "Error: package name is required")
				os.Exit(1)
			}

			if repository == "" {
				errorColor.Fprintln(os.Stderr, "Error: repository URL is required")
				os.Exit(1)
			}

			c := createClient()

			// 检查是否设置了凭据
			if username == "" || apiToken == "" {
				errorColor.Fprintln(os.Stderr, "Error: username and API token are required")
				os.Exit(1)
			}

			infoColor.Fprintf(os.Stderr, "Editing package %s with repository %s...\n", packageName, repository)
			result, err := c.EditPackage(cmd.Context(), packageName, &domain.PackageEditRequest{
				Repository: repository,
			})
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			successColor.Fprintln(os.Stderr, "Successfully edited package!")
			printResult(result)
		},
	}
	editCmd.Flags().StringVar(&packageName, "package", "", "Package name (required)")
	editCmd.Flags().StringVar(&repository, "repository", "", "Repository URL (required)")
	editCmd.MarkFlagRequired("package")
	editCmd.MarkFlagRequired("repository")

	// 子命令: manage update
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update a package",
		Long:  `Update a package in Packagist (requires API credentials).`,
		Run: func(cmd *cobra.Command, args []string) {
			if packageName == "" {
				errorColor.Fprintln(os.Stderr, "Error: package name is required")
				os.Exit(1)
			}

			c := createClient()

			// 检查是否设置了凭据
			if username == "" || apiToken == "" {
				errorColor.Fprintln(os.Stderr, "Error: username and API token are required")
				os.Exit(1)
			}

			infoColor.Fprintf(os.Stderr, "Updating package %s...\n", packageName)
			result, err := c.UpdatePackage(cmd.Context(), &domain.PackageUpdateRequest{
				Repository: packageName,
			})
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			successColor.Fprintln(os.Stderr, "Successfully updated package!")
			printResult(result)
		},
	}
	updateCmd.Flags().StringVar(&packageName, "package", "", "Package name (required)")
	updateCmd.MarkFlagRequired("package")

	manageCmd.AddCommand(createCmd, editCmd, updateCmd)
	return manageCmd
}
