package composer

import "strconv"

// VersionConstraint 表示版本约束类型
type VersionConstraint string

// 预定义的版本约束类型
const (
	// ExactVersion 精确匹配版本，如 1.2.3
	ExactVersion VersionConstraint = "exact"
	// CaretVersion 插入符版本约束，如 ^1.2.3 = >=1.2.3 <2.0.0
	CaretVersion VersionConstraint = "caret"
	// TildeVersion 波浪线版本约束，如 ~1.2.3 = >=1.2.3 <1.3.0
	TildeVersion VersionConstraint = "tilde"
	// RangeVersion 范围版本约束，如 >=1.0 <2.0
	RangeVersion VersionConstraint = "range"
	// WildcardVersion 通配符版本约束，如 1.2.*
	WildcardVersion VersionConstraint = "wildcard"
)

// FormatVersionConstraint 根据约束类型格式化版本字符串
func FormatVersionConstraint(version string, constraintType VersionConstraint) string {
	switch constraintType {
	case ExactVersion:
		return version
	case CaretVersion:
		return "^" + version
	case TildeVersion:
		return "~" + version
	case WildcardVersion:
		// 假设这是一个像 1.2 这样的版本号，转换为 1.2.*
		return version + ".*"
	case RangeVersion:
		// 假设这是一个像 1.2 这样的版本号，转换为 >=1.2.0 <2.0.0
		return ">=" + version + ".0 <" + incrementMajorVersion(version) + ".0.0"
	default:
		return version
	}
}

// incrementMajorVersion 增加主版本号
// 例如：1.2 -> 2，1.2.3 -> 2
func incrementMajorVersion(version string) string {
	// 这里简化处理，只提取第一个数字并增加1
	// 实际应用中应该使用更健壮的版本解析
	for i, c := range version {
		if c >= '0' && c <= '9' {
			major := 0
			for ; i < len(version) && version[i] >= '0' && version[i] <= '9'; i++ {
				major = major*10 + int(version[i]-'0')
			}
			return strconv.Itoa(major + 1)
		}
	}
	return "1" // 如果无法解析，默认返回1
}

// UpdatePackageVersion 更新特定包的版本约束
func (c *Composer) UpdatePackageVersion(packageName string, version string, constraintType VersionConstraint) error {
	formattedVersion := FormatVersionConstraint(version, constraintType)
	// 使用require命令更新现有包的版本
	return c.RequirePackage(packageName, formattedVersion, false)
}

// LockPackageVersion 锁定特定包为确切版本
func (c *Composer) LockPackageVersion(packageName string, version string) error {
	return c.UpdatePackageVersion(packageName, version, ExactVersion)
}

// GetPackageVersions 获取包的可用版本列表
func (c *Composer) GetPackageVersions(packageName string) (string, error) {
	return c.Run("show", "--all", packageName)
}
