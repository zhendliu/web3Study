package model

/*
var DB *gorm.DB

func InitDB() {
	if err := RegisterSysDB(conf.DBConfig.Username, conf.DBConfig.Password, conf.DBConfig.Host, conf.DBConfig.Port, conf.DBConfig.Dbname); err != nil {
		panic(err)
	}
}

func RegisterSysDB(username string, pwd string, host string, port string, dbname string) error {

	DBClient, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&allowOldPasswords=true&parseTime=True&loc=Local", username, pwd, host, port, dbname))
	if err != nil {
		panic(err)
	}

	DBClient.SetMaxIdleConns(2)
	DBClient.SetMaxOpenConns(5)
	DBClient.SetConnMaxLifetime(time.Hour)
	//var dbLog ormLog

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		//dbLog,
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)
	cfg := &gorm.Config{
		Logger: newLogger,
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, pwd, host, port, dbname)

	DB, err = gorm.Open(mysql.Open(dsn), cfg)
	if err != nil {
		return err
	}

	_ = DB.AutoMigrate(&User{})
	_ = DB.AutoMigrate(&Org{})
	_ = DB.AutoMigrate(&Version{})
	_ = DB.AutoMigrate(&Requirement{})
	_ = DB.AutoMigrate(&ReviewDoc{})
	_ = DB.AutoMigrate(&Bug{})
	_ = DB.AutoMigrate(&SmallVersion{})
	_ = DB.AutoMigrate(&Role{})
	_ = DB.AutoMigrate(&Team{})

	_ = DB.AutoMigrate(&Directory{})
	_ = DB.AutoMigrate(&File{})
	_ = DB.AutoMigrate(&Project{})

	_ = DB.AutoMigrate(&ProjectBug{})
	_ = DB.AutoMigrate(&ProjectCustom{})
	_ = DB.AutoMigrate(&OtherTask{})

	_ = DB.AutoMigrate(&Train{})
	_ = DB.AutoMigrate(&Reward{})
	//DB.AutoMigrate(&TrainDoc{})
	_ = DB.AutoMigrate(&PersonGroupChange{})
	_ = DB.AutoMigrate(&Interaction{})

	return nil
}
*/
