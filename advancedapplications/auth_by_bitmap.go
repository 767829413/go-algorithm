package advancedapplications

const (
	// PermNone 没有权限
	PermNone Permission = 0
	// PermGetUserBaiscInfo 获取你的头像、性别、昵称等基本用户信息
	PermGetUserBaiscInfo Permission = 1 << (iota - 1)
	// PermPostWeibo 以你的身份发布微博
	PermPostWeibo
	// PermGetYourFriends 获取你的好友列表
	PermGetYourFriends
	// PermGetYourTimeline 获取你的朋友圈信息
	PermGetYourTimeline
)

// Permission 权限
type Permission uint64

// AddPermission 设置权限
func AddPermission(perm, toAdd Permission) Permission {
	return perm | toAdd
}

// RemovePermission 移除权限
func RemovePermission(perm, toRemove Permission) Permission {
	return perm & (^toRemove)
}
