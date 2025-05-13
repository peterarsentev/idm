package repositories

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"idm/inner/database"
	"idm/inner/models"
	repository "idm/inner/repositories"
	"log"
	"testing"
	"time"
)

func TestCreateKey(t *testing.T) {
	a := assert.New(t)
	db, err := database.NewDb(".env.test")
	if err != nil {
		log.Fatal(err)
	}
	if clearErr := clearDb(db); clearErr != nil {
		log.Printf("Error while truncating tables: %v", clearErr)
	}
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
		if db == nil {
			return
		}
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()
	repo := repository.NewKeyRepository(db)
	filterID, err := NewFixtureFilter(repository.NewFilterRepository(db)).NewFilter("Empty")
	if err != nil {
		log.Fatal(err)
	}
	key := models.Key{
		Name:     "Test Key",
		FilterID: filterID,
	}
	ID, err := repo.Add(context.Background(), key)
	a.NoError(err, "Failed to add key")
	createdKey, err := repo.FindByID(context.Background(), ID)
	a.NoError(err, "Failed to retrieve key")
	a.Equal(key.Name, createdKey.Name, "Key names should match")
}

func TestFindAllKeys(t *testing.T) {
	a := assert.New(t)
	db, err := database.NewDb(".env.test")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if db == nil {
			return
		}
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()
	repo := repository.NewKeyRepository(db)
	keys, err := repo.FindAll(context.Background())
	a.NoError(err, "Failed to fetch keys")
	a.NotEmpty(keys, "Keys should not be empty")
}

func TestCreateKeyValue(t *testing.T) {
	a := assert.New(t)
	db, err := database.NewDb(".env.test")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if db == nil {
			return
		}
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()
	repo := repository.NewKeyValueRepository(db)
	filterID, err := NewFixtureFilter(repository.NewFilterRepository(db)).NewFilter("Empty")
	if err != nil {
		log.Fatal(err)
	}
	keyID, err := NewKeyRepository(repository.NewKeyRepository(db)).NewKey(filterID, "Kye")
	keyValue := models.KeyValue{
		Value: "Test Value",
		KeyID: keyID,
	}
	ID, err := repo.Add(context.Background(), keyValue)
	a.NoError(err, "Failed to add key value")
	createdKeyValue, err := repo.FindByID(context.Background(), ID)
	a.NoError(err, "Failed to retrieve key value")
	a.Equal(keyValue.Value, createdKeyValue.Value, "Key values should match")
}

func TestFindAllKeyValues(t *testing.T) {
	a := assert.New(t)
	db, err := database.NewDb(".env.test")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if db == nil {
			return
		}
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()
	repo := repository.NewKeyValueRepository(db)
	keyValues, err := repo.FindAll(context.Background())
	a.NoError(err, "Failed to fetch key values")
	a.NotEmpty(keyValues, "Key values should not be empty")
}

func TestCreateSpam(t *testing.T) {
	a := assert.New(t)
	db, err := database.NewDb(".env.test")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if db == nil {
			return
		}
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()

	// Create a new chat
	chatRepo := repository.NewChatRepository(db)
	chat := models.Chat{
		ChatID: time.Now().UnixMilli(),
	}
	chatID, err := chatRepo.Add(context.Background(), chat)
	a.NoError(err, "Failed to add chat")

	// Create a new contact
	contactRepo := repository.NewContactRepository(db)
	contact := models.Contact{
		TgUserID: time.Now().UnixMilli(),
		Username: fmt.Sprintf("username_%d", time.Now().UnixMilli()),
		Ham:      0,
		Spam:     1,
		Restrict: false,
	}
	contactID, err := contactRepo.Add(context.Background(), contact)
	a.NoError(err, "Failed to add contact")

	// Create a new spam using the chat and contact created above
	spamRepo := repository.NewSpamRepository(db)
	spam := models.Spam{
		Text:      "Test Spam",
		Time:      time.Now().Format(time.RFC3339),
		ChatID:    chatID,
		ContactID: contactID,
	}
	ID, err := spamRepo.Add(context.Background(), spam)
	a.NoError(err, "Failed to add spam")

	// Retrieve the created spam and validate
	createdSpam, err := spamRepo.FindByID(context.Background(), ID)
	a.NoError(err, "Failed to retrieve spam")
	a.Equal(spam.Text, createdSpam.Text, "Spam texts should match")
}

func TestFindAllSpams(t *testing.T) {
	a := assert.New(t)
	db, err := database.NewDb(".env.test")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if db == nil {
			return
		}
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()
	repo := repository.NewSpamRepository(db)
	spams, err := repo.FindAll(context.Background())
	a.NoError(err, "Failed to fetch spams")
	a.NotEmpty(spams, "Spams should not be empty")
}

func TestCreateVote(t *testing.T) {
	a := assert.New(t)
	db, err := database.NewDb(".env.test")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if db == nil {
			return
		}
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()

	// Create a new User
	roleID, err := repository.NewRoleRepository(db).Add(context.Background(), models.Role{Name: "Role"})
	a.NoError(err, "Failed to add role")
	userRepo := repository.NewUserRepository(db)
	user := models.User{
		Username: fmt.Sprintf("username_%d", time.Now().UnixMilli()),
		Password: "password123",
		Enabled:  true,
		RoleID:   roleID,
	}
	userID, err := userRepo.Add(context.Background(), user)
	a.NoError(err, "Failed to add user")

	// Create a vote using the created user and message
	repo := repository.NewVoteRepository(db)
	vote := models.Vote{
		ChatID:    time.Now().UnixMilli(),
		MessageID: time.Now().UnixMilli(),
		UserID:    userID,
		Vote:      1,
	}
	ID, err := repo.Add(context.Background(), vote)
	a.NoError(err, "Failed to add vote")

	// Retrieve the created vote and validate
	createdVote, err := repo.FindByID(context.Background(), ID)
	a.NoError(err, "Failed to retrieve vote")
	a.Equal(vote.Vote, createdVote.Vote, "Vote values should match")
}

func TestFindAllVotes(t *testing.T) {
	a := assert.New(t)
	db, err := database.NewDb(".env.test")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if db == nil {
			return
		}
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()
	repo := repository.NewVoteRepository(db)
	votes, err := repo.FindAll(context.Background())
	a.NoError(err, "Failed to fetch votes")
	a.NotEmpty(votes, "Votes should not be empty")
}

func TestCreateRole(t *testing.T) {
	a := assert.New(t)
	db, err := database.NewDb(".env.test")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if db == nil {
			return
		}
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()
	repo := repository.NewRoleRepository(db)
	role := models.Role{
		Name: "Test Role",
	}
	ID, err := repo.Add(context.Background(), role)
	a.NoError(err, "Failed to add role")
	createdRole, err := repo.FindByID(context.Background(), ID)
	a.NoError(err, "Failed to retrieve role")
	a.Equal(role.Name, createdRole.Name, "Role names should match")
}

func TestFindAllRoles(t *testing.T) {
	a := assert.New(t)
	db, err := database.NewDb(".env.test")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if db == nil {
			return
		}
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()
	repo := repository.NewRoleRepository(db)
	roles, err := repo.FindAll(context.Background())
	a.NoError(err, "Failed to fetch roles")
	a.NotEmpty(roles, "Roles should not be empty")
}

func TestCreateUser(t *testing.T) {
	a := assert.New(t)
	db, err := database.NewDb(".env.test")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if db == nil {
			return
		}
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()
	role := models.Role{
		Name: "Test Role",
	}
	roleID, err := repository.NewRoleRepository(db).Add(context.Background(), role)
	a.NoError(err, "Failed to add role")
	repo := repository.NewUserRepository(db)
	user := models.User{
		Username: fmt.Sprintf("username_%d", time.Now().UnixMilli()),
		Password: "password123",
		Enabled:  true,
		RoleID:   roleID,
	}
	ID, err := repo.Add(context.Background(), user)
	a.NoError(err, "Failed to add user")
	createdUser, err := repo.FindByID(context.Background(), ID)
	a.NoError(err, "Failed to retrieve user")
	a.Equal(user.Username, createdUser.Username, "Usernames should match")
}

func TestFindAllUsers(t *testing.T) {
	a := assert.New(t)
	db, err := database.NewDb(".env.test")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if db == nil {
			return
		}
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()
	repo := repository.NewUserRepository(db)
	users, err := repo.FindAll(context.Background())
	a.NoError(err, "Failed to fetch users")
	a.NotEmpty(users, "Users should not be empty")
}

func TestCreateContact(t *testing.T) {
	a := assert.New(t)
	db, err := database.NewDb(".env.test")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if db == nil {
			return
		}
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()
	repo := repository.NewContactRepository(db)
	contact := models.Contact{
		TgUserID: time.Now().UnixMilli(),
		Username: fmt.Sprintf("username_%d", time.Now().UnixMilli()),
		Ham:      0,
		Spam:     1,
		Restrict: false,
	}
	ID, err := repo.Add(context.Background(), contact)
	a.NoError(err, "Failed to add contact")
	createdContact, err := repo.FindByID(context.Background(), ID)
	a.NoError(err, "Failed to retrieve contact")
	a.Equal(contact.Username, createdContact.Username, "Usernames should match")
}

func TestFindAllContacts(t *testing.T) {
	a := assert.New(t)
	db, err := database.NewDb(".env.test")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if db == nil {
			return
		}
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()
	repo := repository.NewContactRepository(db)
	contacts, err := repo.FindAll(context.Background())
	a.NoError(err, "Failed to fetch contacts")
	a.NotEmpty(contacts, "Contacts should not be empty")
}

func TestCreateChat(t *testing.T) {
	a := assert.New(t)
	db, err := database.NewDb(".env.test")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if db == nil {
			return
		}
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()
	repo := repository.NewChatRepository(db)
	chat := models.Chat{
		ChatID: time.Now().UnixMilli(),
	}
	ID, err := repo.Add(context.Background(), chat)
	a.NoError(err, "Failed to add chat")
	createdChat, err := repo.FindByID(context.Background(), ID)
	a.NoError(err, "Failed to retrieve chat")
	a.Equal(chat.Title, createdChat.Title, "Chat titles should match")
}

func TestFindAllChats(t *testing.T) {
	a := assert.New(t)
	db, err := database.NewDb(".env.test")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if db == nil {
			return
		}
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()
	repo := repository.NewChatRepository(db)
	chats, err := repo.FindAll(context.Background())
	a.NoError(err, "Failed to fetch chats")
	a.NotEmpty(chats, "Chats should not be empty")
}
