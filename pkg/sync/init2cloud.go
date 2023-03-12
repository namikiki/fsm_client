package sync

//func (s *Syncer) ScannerDirToUploadCloud(rooPath, syncID string) error {
//	rootPathLen := len(rooPath)
//	return filepath.WalkDir(rooPath, func(path string, d fs.DirEntry, err error) error {
//		rawPath := path
//		path = path[rootPathLen:]
//		if d.IsDir() {
//
//			dir := ent.Dir{
//				SyncID:     syncID,
//				Dir:        path + "/", // todo Split
//				Level:      uint64(len(strings.Split(path, "/"))),
//				Deleted:    false,
//				CreateTime: time.Now(),
//				ModTime:    time.Now(),
//			}
//			if err := s.CreateDir(&dir); err != nil {
//				return err
//			}
//			s.db.Create(&dir)
//			return err
//		}
//
//		level := len(strings.Split(path, "/"))
//		suffix := strings.TrimSuffix(path, d.Name())
//		var dir ent.Dir
//		s.db.Where("dir= ? and level = ?", suffix, level-1).Find(&dir)
//
//		info, _ := d.Info()
//
//		file := ent.File{
//			SyncID:      syncID,
//			Name:        info.Name(),
//			ParentDirID: dir.ID,
//			Level:       uint64(level),
//			Hash:        "",
//			Size:        info.Size(),
//			Deleted:     false,
//			CreateTime:  time.Now(),
//			ModTime:     info.ModTime(),
//		}
//
//		open, err := os.Open(rawPath)
//
//		if err := s.CreateFile(&file, open); err != nil {
//			return err
//		}
//		s.db.Create(&file)
//		return err
//
//	})
//
//}
//
//func (s *Syncer) GetSyncTaskToDownload(syncID, path string) error {
//
//	dirs, err := s.GetAllDirBySyncID(syncID)
//	if err != nil {
//		return err
//	}
//
//	for i := range dirs {
//		if err := os.MkdirAll(path+dirs[i].Dir, os.ModePerm); err != nil {
//			if !os.IsExist(err) {
//				log.Println(err)
//				return err
//			}
//		} // 文件夹创建成功
//	}
//	s.db.Create(&dirs)
//
//	files, err := s.GetAllFileBySyncID(syncID)
//	if err != nil {
//		return err
//	}
//
//	for i := range files {
//
//		if fileio, err := s.GetFile(files[i].ID); err == nil {
//			var dir ent.Dir
//			s.db.Where("id = ?", files[i].ParentDirID).Find(&dir)
//			if file, err := os.Create(path + dir.Dir + files[i].Name); err == nil {
//				io.Copy(file, fileio)
//				file.Close()
//			}
//			fileio.Close()
//		}
//
//		s.db.Create(&files)
//	} //文件创建成功
//
//	return nil
//}
//
//func (s *Syncer) CreateSync(name, root string) error {
//	task := ent.SyncTask{
//		UserID:     "xyn233",
//		Type:       "sync",
//		Name:       name,
//		RootDir:    root,
//		Deleted:    false,
//		CreateTime: time.Now(),
//	}
//
//	if err := s.CreateSyncTask(&task); err != nil {
//		return err
//	}
//	s.db.Create(&task)
//
//	if err := s.ScannerDirToUploadCloud(task.RootDir, task.ID); err != nil {
//		return err
//	}
//	return nil
//}
//
//type PubSubMessage struct {
//	Type     string
//	Action   string
//	ClientID string
//	Data     interface{}
//}
//
//func (s *Syncer) WebSocketConn() {
//	wsDialer := websocket.Dialer{}
//
//	// 使用 HTTP 客户端与 WebSocket 服务器建立连接
//	wsConn, _, err := wsDialer.Dial("ws://localhost:8080/websocketconn?uid=xyn233&cid="+s.clientID, http.Header{})
//	if err != nil {
//		log.Fatal("连接失败：", err)
//	}
//
//	for {
//		messageType, receivedMessage, err := wsConn.ReadMessage()
//		if err != nil {
//			log.Fatal("接收消息失败：", err)
//		}
//
//		var psm PubSubMessage
//		if err := json.Unmarshal(receivedMessage, &psm); err != nil {
//			log.Println(err)
//		}
//
//		if psm.ClientID == s.clientID {
//			continue
//		}
//
//		fmt.Printf("接收到的消息类型：%d\n", messageType)
//		fmt.Printf("接收到的消息内容：%s\n", string(receivedMessage))
//	}
//
//}
