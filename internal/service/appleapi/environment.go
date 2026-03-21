package appleapi

// Environment 表示App Store服务器API的环境
type Environment string

const (
	// Production 生产环境
	Production Environment = "PRODUCTION"
	// Sandbox 沙盒环境
	Sandbox Environment = "SANDBOX"
)

// Valid 检查环境是否有效
func (e Environment) Valid() bool {
	return e == Production || e == Sandbox
}

// BaseURL 获取环境对应的基础URL
func (e Environment) BaseURL() string {
	switch e {
	case Production:
		return "https://api.storekit.itunes.apple.com/inApps"
	case Sandbox:
		return "https://api.storekit-sandbox.itunes.apple.com/inApps"
	default:
		return ""
	}
}
