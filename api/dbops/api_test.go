package dbops

import (
	"testing"
	// "github.com/gitcloneese/video_server/api/defs"
	"fmt"
	"strconv"
	"time"

	"github.com/gitcloneese/video_server/api/utils"
)

var (
	tempvid string
	tempsid string
	tempcid string
)

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate, comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()

}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("ReGet", testRegetUser)

}

func testAddUser(t *testing.T) {
	err := AddUserCredential("avenssi", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("avenssi")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser")
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("avenssi", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}

}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("avenssi")
	if err != nil {
		t.Errorf("Error of ReGetUser: %v", err)
	}
	if pwd != "" {
		t.Errorf("Deleting user test Failed")
	}

}

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}
	tempvid = vi.Id
}

func testGetVideoInfo(t *testing.T) {
	video, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}

	if video.AuthorId != 1 || video.Name != "my-video" {
		t.Error("Error of GetVideoInfo 2")
	}

}

func testDeleteVideoInfo(t *testing.T) {
	err := DelteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi != nil {
		t.Errorf("Error of RegetVideoInfo: %v", err)
	}
}

func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddComments", testAddComment)
	t.Run("ListComments", testListComments)
	t.Run("DeleteComments", testDeleteComments)
	clearTables()
}

func testAddComment(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "I like this video"
	cid, err := AddNewComments(vid, aid, content)
	if err != nil {
		t.Errorf("Error of AddConmments: %v", err)
	}
	tempcid = cid
}

func testDeleteComments(t *testing.T) {
	vid := "12345"
	err := DeleteComments(vid, tempcid)
	if err != nil {
		t.Errorf("Error of DeleteComments: %v", err)
	}
}

func testListComments(t *testing.T) {
	vid := "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))

	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}

	for i, element := range res {
		fmt.Printf("Comments : %d, %v \n", i, element)
	}

}

func TestSession(t *testing.T) {
	clearTables()
	t.Run("TestAddSession", testAddSession)
	t.Run("TestRetrieveSession", testRetrieveSession)
	t.Run("testDelteSession", testDeleteSession)
	t.Run("testreRetrieveSession", testreRetrieveSession)
	clearTables()
}

func testAddSession(t *testing.T) {
	sid, err := utils.NewUUID()
	if err != nil {
		t.Errorf("Error of UUID, %v", err)
	}

	tempsid = sid
	ttl := int64(129183174987124)
	err = InsertSession(sid, ttl, "skyone")
	if err != nil {
		t.Errorf("Error of InsertSession: %v", err)
	}
}

func testRetrieveSession(t *testing.T) {
	res, err := RetriveSession(tempsid)
	if err != nil {
		t.Errorf("Error of RetrieveSession :%v", err)
	}

	fmt.Printf("session :%v", res)
}

func testDeleteSession(t *testing.T) {
	DeleteSession(tempsid)
}

func testreRetrieveSession(t *testing.T) {
	res, err := RetriveSession(tempsid)
	if err != nil || res != nil {
		t.Errorf("Error of reRetrieveSession:%v ", err)
	}

	fmt.Printf("session :%v", res)
}
