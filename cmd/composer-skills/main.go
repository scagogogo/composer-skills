package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/scagogogo/composer-skills/pkg/client"
	"github.com/scagogogo/composer-skills/pkg/composer"
	"github.com/scagogogo/composer-skills/pkg/domain"
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
		Use:   "composer-skills",
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
	rootCmd.AddCommand(createLocalCmd())

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

// createLocalCmd creates the local Composer CLI wrapper commands
func createLocalCmd() *cobra.Command {
	localCmd := &cobra.Command{
		Use:   "local",
		Short: "Local Composer CLI operations",
		Long:  `Commands for managing PHP projects using the local Composer binary.`,
	}

	var workingDir string
	var noDev bool
	var optimize bool

	// local install
	installCmd := &cobra.Command{
		Use:   "install",
		Short: "Install dependencies",
		Long:  `Install project dependencies using composer install.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Installing dependencies...")
			if err := comp.Install(noDev, optimize); err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintln(os.Stderr, "Dependencies installed successfully!")
		},
	}
	installCmd.Flags().BoolVar(&noDev, "no-dev", false, "Skip installing dev dependencies")
	installCmd.Flags().BoolVar(&optimize, "optimize", false, "Optimize autoloader")

	// local require
	var version string
	requireCmd := &cobra.Command{
		Use:   "require [package]",
		Short: "Add a package dependency",
		Long:  `Add a new package dependency using composer require.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintf(os.Stderr, "Adding package %s...\n", args[0])
			if err := comp.RequirePackage(args[0], version, noDev); err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintf(os.Stderr, "Package %s added successfully!\n", args[0])
		},
	}
	requireCmd.Flags().BoolVar(&noDev, "dev", false, "Add as dev dependency")
	requireCmd.Flags().StringVar(&version, "version", "", "Version constraint (e.g. ^5.0)")

	// local update
	updateCmd := &cobra.Command{
		Use:   "update [packages...]",
		Short: "Update dependencies",
		Long:  `Update project dependencies using composer update.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Updating dependencies...")
			if err := comp.Update(args, noDev); err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintln(os.Stderr, "Dependencies updated successfully!")
		},
	}
	updateCmd.Flags().BoolVar(&noDev, "no-dev", false, "Skip updating dev dependencies")

	// local remove
	removeCmd := &cobra.Command{
		Use:   "remove [package]",
		Short: "Remove a package dependency",
		Long:  `Remove a package dependency using composer remove.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintf(os.Stderr, "Removing package %s...\n", args[0])
			if err := comp.Remove(args[0], noDev); err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintf(os.Stderr, "Package %s removed successfully!\n", args[0])
		},
	}
	removeCmd.Flags().BoolVar(&noDev, "dev", false, "Remove from dev dependencies")

	// local audit
	auditCmd := &cobra.Command{
		Use:   "audit",
		Short: "Run security audit",
		Long:  `Run a security audit on project dependencies.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Running security audit...")
			result, err := comp.AuditWithJSON()
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintf(os.Stderr, "Audit complete. Found %d vulnerabilities.\n", result.Found)
			printResult(result)
		},
	}

	// local version
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show Composer version",
		Long:  `Show the installed Composer version.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			ver, err := comp.GetVersion()
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintf(os.Stderr, "Composer version: %s\n", ver)
		},
	}

	// local validate
	validateCmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate composer.json",
		Long:  `Validate the composer.json file.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Validating composer.json...")
			if err := comp.Validate(); err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintln(os.Stderr, "composer.json is valid!")
		},
	}

	// local outdated
	outdatedCmd := &cobra.Command{
		Use:   "outdated",
		Short: "Show outdated packages",
		Long:  `Show outdated packages in the project.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Checking for outdated packages...")
			output, err := comp.OutdatedPackages()
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(output)
		},
	}

	// local show
	showCmd := &cobra.Command{
		Use:   "show [package]",
		Short: "Show package information",
		Long:  `Show information about installed packages.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			if len(args) > 0 {
				output, err := comp.ShowPackage(args[0])
				if err != nil {
					errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
					os.Exit(1)
				}
				fmt.Println(output)
			} else {
				output, err := comp.ShowAllPackages()
				if err != nil {
					errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
					os.Exit(1)
				}
				fmt.Println(output)
			}
		},
	}

	// local create-project
	var projectVersion string
	createProjectCmd := &cobra.Command{
		Use:   "create-project [package] [directory]",
		Short: "Create a new project",
		Long:  `Create a new PHP project from a package.`,
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintf(os.Stderr, "Creating project %s in %s...\n", args[0], args[1])
			if err := comp.CreateProject(args[0], args[1], projectVersion); err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintln(os.Stderr, "Project created successfully!")
		},
	}
	createProjectCmd.Flags().StringVar(&projectVersion, "version", "", "Package version")

	// ========== 新增子命令 ==========

	// local dump-autoload
	dumpAutoloadCmd := &cobra.Command{
		Use:   "dump-autoload",
		Short: "Generate autoloader files",
		Long:  `Generate Composer autoloader files, optionally with optimization.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Generating autoloader files...")
			if err := comp.DumpAutoload(optimize); err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintln(os.Stderr, "Autoloader files generated successfully!")
		},
	}
	dumpAutoloadCmd.Flags().BoolVar(&optimize, "optimize", false, "Optimize autoloader")

	// local init
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a new Composer project",
		Long:  `Initialize a new Composer project by creating a composer.json file.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Initializing new Composer project...")
			if err := comp.InitProject(); err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintln(os.Stderr, "Composer project initialized successfully!")
		},
	}

	// local init-with-options
	var initName, initDescription, initAuthor string
	initWithOptionsCmd := &cobra.Command{
		Use:   "init-with-options",
		Short: "Initialize a new Composer project with options",
		Long:  `Initialize a new Composer project with name, description, and author options.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Initializing new Composer project with options...")
			options := map[string]string{
				"no-interaction": "",
			}
			if err := comp.InitProjectWithOptions(initName, initDescription, initAuthor, options); err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintln(os.Stderr, "Composer project initialized successfully!")
		},
	}
	initWithOptionsCmd.Flags().StringVar(&initName, "name", "", "Project name (vendor/name format)")
	initWithOptionsCmd.Flags().StringVar(&initDescription, "description", "", "Project description")
	initWithOptionsCmd.Flags().StringVar(&initAuthor, "author", "", "Author info (Name <email>)")

	// local run-script
	runScriptCmd := &cobra.Command{
		Use:   "run-script [script-name] [args...]",
		Short: "Run a Composer script",
		Long:  `Run a script defined in composer.json, optionally passing additional arguments.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			scriptName := args[0]
			scriptArgs := args[1:]
			infoColor.Fprintf(os.Stderr, "Running script %s...\n", scriptName)
			output, err := comp.RunScript(scriptName, scriptArgs...)
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(output)
			successColor.Fprintf(os.Stderr, "Script %s completed successfully!\n", scriptName)
		},
	}

	// local list-scripts
	listScriptsCmd := &cobra.Command{
		Use:   "list-scripts",
		Short: "List all defined scripts",
		Long:  `List all scripts defined in composer.json.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Listing scripts...")
			output, err := comp.ListScripts()
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(output)
		},
	}

	// local search
	localSearchCmd := &cobra.Command{
		Use:   "search [query]",
		Short: "Search for packages locally",
		Long:  `Search for packages using the local Composer binary.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintf(os.Stderr, "Searching for packages matching '%s'...\n", args[0])
			output, err := comp.Search(args[0])
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(output)
		},
	}

	// local depends
	dependsCmd := &cobra.Command{
		Use:   "depends [package]",
		Short: "Show reverse dependencies",
		Long:  `Show which packages depend on the specified package.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintf(os.Stderr, "Checking reverse dependencies for %s...\n", args[0])
			output, err := comp.ShowReverseDependencies(args[0])
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(output)
		},
	}

	// local why
	whyCmd := &cobra.Command{
		Use:   "why [package]",
		Short: "Explain why a package is installed",
		Long:  `Explain why a package is installed by showing which packages require it.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintf(os.Stderr, "Checking why %s is installed...\n", args[0])
			output, err := comp.WhyPackage(args[0])
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(output)
		},
	}

	// local why-not
	whyNotCmd := &cobra.Command{
		Use:   "why-not [package] [version]",
		Short: "Explain why a package version cannot be installed",
		Long:  `Explain why a specific version of a package cannot be installed.`,
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintf(os.Stderr, "Checking why %s version %s cannot be installed...\n", args[0], args[1])
			output, err := comp.WhyNotPackage(args[0], args[1])
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(output)
		},
	}

	// local tree
	treeCmd := &cobra.Command{
		Use:   "tree [package]",
		Short: "Show dependency tree",
		Long:  `Show the dependency tree for a package or the entire project.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			pkgName := ""
			if len(args) > 0 {
				pkgName = args[0]
			}
			infoColor.Fprintln(os.Stderr, "Displaying dependency tree...")
			output, err := comp.ShowDependencyTree(pkgName)
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(output)
		},
	}

	// local reinstall
	reinstallCmd := &cobra.Command{
		Use:   "reinstall [package]",
		Short: "Reinstall a package",
		Long:  `Reinstall a package by removing and requiring it again.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintf(os.Stderr, "Reinstalling package %s...\n", args[0])
			if err := comp.Reinstall(args[0]); err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintf(os.Stderr, "Package %s reinstalled successfully!\n", args[0])
		},
	}

	// local bump
	var bumpDevOnly, bumpPreferStable, bumpDryRun bool
	bumpCmd := &cobra.Command{
		Use:   "bump [packages...]",
		Short: "Bump packages to latest compatible versions",
		Long:  `Bump packages to their latest compatible versions within composer.json constraints.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintf(os.Stderr, "Bumping packages...\n")
			options := map[string]string{}
			if bumpDevOnly {
				options["dev-only"] = ""
			}
			if bumpPreferStable {
				options["prefer-stable"] = ""
			}
			if bumpDryRun {
				options["dry-run"] = ""
			}
			if len(options) > 0 {
				if err := comp.BumpPackagesWithOptions(args, options); err != nil {
					errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
					os.Exit(1)
				}
			} else {
				if err := comp.BumpPackages(args); err != nil {
					errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
					os.Exit(1)
				}
			}
			successColor.Fprintln(os.Stderr, "Packages bumped successfully!")
		},
	}
	bumpCmd.Flags().BoolVar(&bumpDevOnly, "dev-only", false, "Only bump dev dependencies")
	bumpCmd.Flags().BoolVar(&bumpPreferStable, "prefer-stable", false, "Prefer stable versions")
	bumpCmd.Flags().BoolVar(&bumpDryRun, "dry-run", false, "Only show what would be bumped")

	// local browse
	browseCmd := &cobra.Command{
		Use:   "browse [package]",
		Short: "Open package homepage in browser",
		Long:  `Open the homepage of a package in the default browser.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintf(os.Stderr, "Opening homepage for %s...\n", args[0])
			if err := comp.BrowsePackage(args[0]); err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintf(os.Stderr, "Opened homepage for %s!\n", args[0])
		},
	}

	// local self-update
	selfUpdateCmd := &cobra.Command{
		Use:   "self-update",
		Short: "Update Composer to the latest version",
		Long:  `Update the Composer binary itself to the latest version.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Updating Composer...")
			if err := comp.SelfUpdate(); err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintln(os.Stderr, "Composer updated successfully!")
		},
	}

	// local check
	checkCmd := &cobra.Command{
		Use:   "check",
		Short: "Check dependency validity",
		Long:  `Check that dependencies are valid and composer.json is in sync with composer.lock.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Checking dependencies...")
			output, err := comp.Check()
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(output)
			successColor.Fprintln(os.Stderr, "Dependency check complete!")
		},
	}

	// local diagnose
	diagnoseCmd := &cobra.Command{
		Use:   "diagnose",
		Short: "Diagnose system for common errors",
		Long:  `Run Composer's diagnostic checks to identify common errors.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Running diagnostics...")
			output, err := comp.Diagnose()
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(output)
		},
	}

	// local status
	statusCmd := &cobra.Command{
		Use:   "status",
		Short: "Show local modifications",
		Long:  `Show local modifications to installed packages.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Checking package status...")
			output, err := comp.Status()
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(output)
		},
	}

	// local fund
	var fundJSON bool
	fundCmd := &cobra.Command{
		Use:   "fund [package]",
		Short: "Show funding information",
		Long:  `Show funding information for packages. Optionally specify a package name.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			if fundJSON {
				infoColor.Fprintln(os.Stderr, "Fetching funding information (JSON)...")
				result, err := comp.FundWithJSON()
				if err != nil {
					errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
					os.Exit(1)
				}
				printResult(result)
			} else if len(args) > 0 {
				infoColor.Fprintf(os.Stderr, "Fetching funding information for %s...\n", args[0])
				output, err := comp.FundWithPackage(args[0])
				if err != nil {
					errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
					os.Exit(1)
				}
				fmt.Println(output)
			} else {
				infoColor.Fprintln(os.Stderr, "Fetching funding information...")
				output, err := comp.Fund()
				if err != nil {
					errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
					os.Exit(1)
				}
				fmt.Println(output)
			}
		},
	}
	fundCmd.Flags().BoolVar(&fundJSON, "json", false, "Output funding info as JSON")

	// local global - 子命令组
	globalCmd := &cobra.Command{
		Use:   "global",
		Short: "Global Composer operations",
		Long:  `Commands for managing globally installed Composer packages.`,
	}

	// local global require
	var globalVersion string
	globalRequireCmd := &cobra.Command{
		Use:   "require [package]",
		Short: "Globally require a package",
		Long:  `Install a package globally using composer global require.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintf(os.Stderr, "Globally requiring package %s...\n", args[0])
			if err := comp.GlobalRequire(args[0], globalVersion); err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintf(os.Stderr, "Package %s globally required successfully!\n", args[0])
		},
	}
	globalRequireCmd.Flags().StringVar(&globalVersion, "version", "", "Version constraint (e.g. ^5.0)")

	// local global update
	globalUpdateCmd := &cobra.Command{
		Use:   "update [packages...]",
		Short: "Globally update packages",
		Long:  `Update globally installed packages using composer global update.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Globally updating packages...")
			if err := comp.GlobalUpdate(args); err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintln(os.Stderr, "Global packages updated successfully!")
		},
	}

	// local global remove
	globalRemoveCmd := &cobra.Command{
		Use:   "remove [package]",
		Short: "Globally remove a package",
		Long:  `Remove a globally installed package using composer global remove.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintf(os.Stderr, "Globally removing package %s...\n", args[0])
			if err := comp.GlobalRemove(args[0]); err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintf(os.Stderr, "Package %s globally removed successfully!\n", args[0])
		},
	}

	// local global install
	globalInstallCmd := &cobra.Command{
		Use:   "install",
		Short: "Globally install dependencies",
		Long:  `Install global dependencies using composer global install.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Globally installing dependencies...")
			if err := comp.GlobalInstall(); err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintln(os.Stderr, "Global dependencies installed successfully!")
		},
	}

	// local global list
	globalListCmd := &cobra.Command{
		Use:   "list",
		Short: "List globally installed packages",
		Long:  `List all globally installed Composer packages.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Listing globally installed packages...")
			output, err := comp.GlobalList()
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(output)
		},
	}

	// 为global子命令添加--working-dir标志
	for _, subCmd := range []*cobra.Command{globalRequireCmd, globalUpdateCmd, globalRemoveCmd, globalInstallCmd, globalListCmd} {
		subCmd.Flags().StringVar(&workingDir, "working-dir", "", "Working directory for Composer operations")
	}

	globalCmd.AddCommand(globalRequireCmd, globalUpdateCmd, globalRemoveCmd, globalInstallCmd, globalListCmd)

	// local licenses
	var licensesFormat string
	licensesCmd := &cobra.Command{
		Use:   "licenses",
		Short: "Show license information",
		Long:  `Show license information for project dependencies.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Fetching license information...")
			var output string
			var err error
			if licensesFormat != "" {
				output, err = comp.LicensesWithFormat(licensesFormat)
			} else {
				output, err = comp.Licenses()
			}
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(output)
		},
	}
	licensesCmd.Flags().StringVar(&licensesFormat, "format", "", "Output format (e.g. json, text)")

	// local check-licenses
	checkLicensesCmd := &cobra.Command{
		Use:   "check-licenses",
		Short: "Check license compatibility",
		Long:  `Check the license compatibility of project dependencies.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Checking license compatibility...")
			output, err := comp.CheckLicenses()
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(output)
		},
	}

	// local clear-cache
	clearCacheCmd := &cobra.Command{
		Use:   "clear-cache",
		Short: "Clear Composer cache",
		Long:  `Clear the Composer cache directory.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Clearing Composer cache...")
			if err := comp.ClearCache(); err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintln(os.Stderr, "Composer cache cleared successfully!")
		},
	}

	// local archive
	var archiveFormat string
	archiveCmd := &cobra.Command{
		Use:   "archive [destination]",
		Short: "Create a project archive",
		Long:  `Create an archive of the current project in the specified destination directory.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintf(os.Stderr, "Creating archive in %s...\n", args[0])
			var output string
			var err error
			if archiveFormat != "" {
				output, err = comp.ArchiveWithFormat(args[0], archiveFormat)
			} else {
				output, err = comp.Archive(args[0])
			}
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(output)
			successColor.Fprintln(os.Stderr, "Archive created successfully!")
		},
	}
	archiveCmd.Flags().StringVar(&archiveFormat, "format", "", "Archive format (zip or tar)")

	// local check-platform
	checkPlatformCmd := &cobra.Command{
		Use:   "check-platform",
		Short: "Check platform requirements",
		Long:  `Check platform requirements and output results as JSON.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Checking platform requirements...")
			result, err := comp.CheckPlatform()
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			printResult(result)
		},
	}

	// local get-php-version
	getPHPVersionCmd := &cobra.Command{
		Use:   "get-php-version",
		Short: "Get PHP version",
		Long:  `Get the current PHP version used by Composer.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Getting PHP version...")
			phpVer, err := comp.GetPHPVersion()
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintf(os.Stderr, "PHP version: %s\n", phpVer)
		},
	}

	// local has-extension
	hasExtensionCmd := &cobra.Command{
		Use:   "has-extension [extension]",
		Short: "Check if a PHP extension is installed",
		Long:  `Check whether a specific PHP extension is installed.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintf(os.Stderr, "Checking if extension %s is installed...\n", args[0])
			has, err := comp.HasExtension(args[0])
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			if has {
				successColor.Fprintf(os.Stderr, "Extension %s is installed.\n", args[0])
			} else {
				warnColor.Fprintf(os.Stderr, "Extension %s is NOT installed.\n", args[0])
			}
		},
	}

	// local exec
	execCmd := &cobra.Command{
		Use:   "exec [binary] [args...]",
		Short: "Execute a vendor binary",
		Long:  `Execute a vendor binary installed by Composer, optionally passing additional arguments.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			binary := args[0]
			execArgs := args[1:]
			infoColor.Fprintf(os.Stderr, "Executing binary %s...\n", binary)
			output, err := comp.Exec(binary, execArgs...)
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(output)
		},
	}

	// local suggest
	suggestCmd := &cobra.Command{
		Use:   "suggest",
		Short: "Show suggested packages",
		Long:  `Show packages suggested by installed dependencies.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Fetching suggested packages...")
			if err := comp.Suggests(); err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
		},
	}

	// local config - 子命令组
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Manage Composer configuration",
		Long:  `Get and set Composer configuration values.`,
	}

	// local config get
	var configGlobal bool
	configGetCmd := &cobra.Command{
		Use:   "get [setting]",
		Short: "Get a Composer config value",
		Long:  `Get a Composer configuration value, optionally from the global config.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintf(os.Stderr, "Getting config value for %s...\n", args[0])
			value, err := comp.GetConfigWithGlobal(args[0], configGlobal)
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(value)
		},
	}
	configGetCmd.Flags().BoolVar(&configGlobal, "global", false, "Read from global config")

	// local config set
	configSetCmd := &cobra.Command{
		Use:   "set [setting] [value]",
		Short: "Set a Composer config value",
		Long:  `Set a Composer configuration value, optionally in the global config.`,
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintf(os.Stderr, "Setting config %s = %s...\n", args[0], args[1])
			if err := comp.SetConfigWithGlobal(args[0], args[1], configGlobal); err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintf(os.Stderr, "Config %s set successfully!\n", args[0])
		},
	}
	configSetCmd.Flags().BoolVar(&configGlobal, "global", false, "Write to global config")

	// 为config子命令添加--working-dir标志
	for _, subCmd := range []*cobra.Command{configGetCmd, configSetCmd} {
		subCmd.Flags().StringVar(&workingDir, "working-dir", "", "Working directory for Composer operations")
	}

	configCmd.AddCommand(configGetCmd, configSetCmd)

	// local auth - 子命令组
	authCmd := &cobra.Command{
		Use:   "auth",
		Short: "Manage Composer authentication",
		Long:  `Commands for managing Composer authentication configuration.`,
	}

	// local auth show
	authShowCmd := &cobra.Command{
		Use:   "show",
		Short: "Show authentication configuration",
		Long:  `Show the current Composer authentication configuration.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Fetching authentication configuration...")
			config, err := comp.GetAuthConfig()
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			printResult(config)
		},
	}

	// local auth add-github
	authAddGitHubCmd := &cobra.Command{
		Use:   "add-github [domain] [token]",
		Short: "Add a GitHub OAuth token",
		Long:  `Add a GitHub OAuth token for the specified domain.`,
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintf(os.Stderr, "Adding GitHub token for %s...\n", args[0])
			if err := comp.AddGitHubToken(args[0], args[1]); err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintln(os.Stderr, "GitHub token added successfully!")
		},
	}

	// local auth add-gitlab
	authAddGitLabCmd := &cobra.Command{
		Use:   "add-gitlab [domain] [token]",
		Short: "Add a GitLab OAuth token",
		Long:  `Add a GitLab OAuth token for the specified domain.`,
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintf(os.Stderr, "Adding GitLab token for %s...\n", args[0])
			if err := comp.AddGitLabToken(args[0], args[1]); err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintln(os.Stderr, "GitLab token added successfully!")
		},
	}

	// local auth add-bearer
	authAddBearerCmd := &cobra.Command{
		Use:   "add-bearer [domain] [token]",
		Short: "Add a Bearer token",
		Long:  `Add a Bearer authentication token for the specified domain.`,
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintf(os.Stderr, "Adding Bearer token for %s...\n", args[0])
			if err := comp.AddBearerToken(args[0], args[1]); err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintln(os.Stderr, "Bearer token added successfully!")
		},
	}

	// local auth add-http-basic
	authAddHTTPBasicCmd := &cobra.Command{
		Use:   "add-http-basic [domain] [username] [password]",
		Short: "Add HTTP Basic authentication",
		Long:  `Add HTTP Basic authentication credentials for the specified domain.`,
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintf(os.Stderr, "Adding HTTP Basic auth for %s...\n", args[0])
			if err := comp.AddHTTPBasicAuth(args[0], args[1], args[2]); err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintln(os.Stderr, "HTTP Basic auth added successfully!")
		},
	}

	// local auth remove
	authRemoveCmd := &cobra.Command{
		Use:   "remove [auth-type] [domain]",
		Short: "Remove an authentication token",
		Long:  `Remove an authentication token of the specified type for the given domain. Auth type can be: github-oauth, gitlab-oauth, bitbucket-oauth, bearer, http-basic.`,
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintf(os.Stderr, "Removing %s token for %s...\n", args[0], args[1])
			if err := comp.RemoveToken(args[0], args[1]); err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintln(os.Stderr, "Token removed successfully!")
		},
	}

	// 为auth子命令添加--working-dir标志
	for _, subCmd := range []*cobra.Command{authShowCmd, authAddGitHubCmd, authAddGitLabCmd, authAddBearerCmd, authAddHTTPBasicCmd, authRemoveCmd} {
		subCmd.Flags().StringVar(&workingDir, "working-dir", "", "Working directory for Composer operations")
	}

	authCmd.AddCommand(authShowCmd, authAddGitHubCmd, authAddGitLabCmd, authAddBearerCmd, authAddHTTPBasicCmd, authRemoveCmd)

	// local home
	homeCmd := &cobra.Command{
		Use:   "home",
		Short: "Get Composer home directory",
		Long:  `Get the Composer home directory path.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Getting Composer home directory...")
			home, err := comp.GetComposerHome()
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Fprintf(os.Stderr, "Composer home: %s\n", home)
		},
	}

	// local validate-lock
	validateLockCmd := &cobra.Command{
		Use:   "validate-lock",
		Short: "Validate composer.lock",
		Long:  `Validate the composer.lock file, checking if it exists and is in sync with composer.json.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Validating composer.lock...")
			output, err := comp.ValidateComposerLock()
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(output)
			successColor.Fprintln(os.Stderr, "composer.lock is valid!")
		},
	}

	// local normalize
	normalizeCmd := &cobra.Command{
		Use:   "normalize",
		Short: "Normalize composer.json",
		Long:  `Normalize the composer.json file to follow best practices formatting.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Normalizing composer.json...")
			output, err := comp.NormalizeComposerJson()
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(output)
			successColor.Fprintln(os.Stderr, "composer.json normalized successfully!")
		},
	}

	// local environment
	environmentCmd := &cobra.Command{
		Use:   "environment",
		Short: "Get Composer environment info",
		Long:  `Get detailed Composer environment information and configuration.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Fetching environment information...")
			info, err := comp.GetEnvironmentInfo()
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			printResult(info)
		},
	}

	// local project-info
	projectInfoCmd := &cobra.Command{
		Use:   "project-info",
		Short: "Get project information",
		Long:  `Get current project information from composer.json.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			infoColor.Fprintln(os.Stderr, "Fetching project information...")
			info, err := comp.GetProjectInfo()
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			printResult(info)
		},
	}

	// local about
	aboutCmd := &cobra.Command{
		Use:   "about",
		Short: "Show short information about Composer",
		Long:  `Display a short description of Composer.`,
		Run: func(cmd *cobra.Command, args []string) {
			comp := createComposerInstance(workingDir)
			output, err := comp.About()
			if err != nil {
				errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(output)
		},
	}

	// Add common --working-dir flag to all subcommands
	allSubCmds := []*cobra.Command{
		installCmd, requireCmd, updateCmd, removeCmd, auditCmd, versionCmd, validateCmd,
		outdatedCmd, showCmd, createProjectCmd,
		dumpAutoloadCmd, initCmd, initWithOptionsCmd, runScriptCmd, listScriptsCmd,
		localSearchCmd, dependsCmd, whyCmd, whyNotCmd, treeCmd, reinstallCmd, bumpCmd,
		browseCmd, selfUpdateCmd, checkCmd, diagnoseCmd, statusCmd, fundCmd,
		globalCmd, licensesCmd, checkLicensesCmd, clearCacheCmd, archiveCmd,
		checkPlatformCmd, getPHPVersionCmd, hasExtensionCmd, execCmd, suggestCmd,
		configCmd, authCmd, homeCmd, validateLockCmd, normalizeCmd, environmentCmd,
		projectInfoCmd,
		aboutCmd,
	}
	for _, subCmd := range allSubCmds {
		// 跳过已经有--working-dir的子命令组（它们的子命令已经单独添加了）
		if subCmd == globalCmd || subCmd == configCmd || subCmd == authCmd {
			continue
		}
		subCmd.Flags().StringVar(&workingDir, "working-dir", "", "Working directory for Composer operations")
	}

	localCmd.AddCommand(allSubCmds...)
	return localCmd
}

// createComposerInstance creates a Composer CLI wrapper instance
func createComposerInstance(workingDir string) *composer.Composer {
	options := composer.DefaultOptions()
	if workingDir != "" {
		options.WorkingDir = workingDir
	}
	comp, err := composer.New(options)
	if err != nil {
		errorColor.Fprintf(os.Stderr, "Error creating Composer instance: %v\n", err)
		os.Exit(1)
	}
	return comp
}
