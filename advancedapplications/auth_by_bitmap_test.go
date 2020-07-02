package advancedapplications

import (
	"testing"
)

func TestAuthByBitMap(t *testing.T) {

	t.Logf("4 basic permissions: %v, %v, %v, %v\n", PermGetUserBaiscInfo, PermPostWeibo, PermGetYourFriends, PermGetYourTimeline)
	p := PermNone
	t.Logf("before set perm: %v\n", p)
	p = AddPermission(p, PermGetUserBaiscInfo)
	t.Logf("after set perm: %v\n", p)
	p = AddPermission(p, PermGetYourTimeline)
	t.Logf("after set perm: %v\n", p)
}
