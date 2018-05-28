package storage

import (
	"strings"
	"testing"
)

func TestCheckQuotaNotifiesUser(t *testing.T) {
	// 保存留待恢复的notifyUser
	saved := notifyUser
	// 测试过后(不管成功或是失败)恢复被测试的方法
	defer func() {
		notifyUser = saved
	}()
	var notifiedUser, notifiedMsg string
	// notifyUser方法根据需要被修改
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}
	// 模拟使用了980MB
	const user = "joe@example.org"
	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}
	if notifiedUser != user {
		t.Errorf("wrong user(%s) notified,want %s", notifiedUser, user)
	}
	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected noptification message <<%s>>,"+"want substring %q", notifiedMsg, wantSubstring)
	}
}
