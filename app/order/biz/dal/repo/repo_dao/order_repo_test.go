package repo_dao

import (
	"gorm.io/gorm/logger"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/All-Done-Right/douyin-mall-microservice/app/order/biz/dal/repo/model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 创建模拟数据库
func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	// 创建sqlmock

	mockDB, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// 使用sqlmock创建gorm.DB实例
	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "repo",
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	})

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}

	return db, mock
}

func TestCreateOrder(t *testing.T) {
	// 定义当前时间作为创建时间
	now := time.Now()

	// 设置测试用例
	testCases := []struct {
		name          string
		order         model.Order
		mockSetup     func(sqlmock.Sqlmock)
		expectedID    string
		expectedError bool
	}{
		{
			name: "成功创建订单",
			order: model.Order{
				Model: gorm.Model{
					ID:        2, // 让数据库自动生成
					CreatedAt: now,
					UpdatedAt: now,
				},
				OrderID:      "order123",
				UserID:       1001,
				UserCurrency: "USD",
				Email:        "test@example.com",
				OrderItems: []model.OrderItem{
					{
						OrderID:   "order123",
						ProductID: 5001,
						Quantity:  2,
						Cost:      99.99,
					},
				},
				Address: model.Address{

					StreetAddress: "123 Main St",
					City:          "New York",
					State:         "NY",
					Country:       "USA",
					ZipCode:       10001,
				},
				Paid:    false,
				Expired: false,
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				// 期望执行主订单插入
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `orders`")).
					WithArgs(
						sqlmock.AnyArg(),   // CreatedAt
						sqlmock.AnyArg(),   // UpdatedAt
						nil,                // DeletedAt
						"order123",         // OrderID
						uint32(1001),       // UserID
						"USD",              // UserCurrency
						"test@example.com", // Address.Email
						"123 Main St",      // StreetAddress
						"New York",         // City
						"NY",               // State
						"USA",              // Country
						int32(10001),       // ZipCode
						false,              // Paid
						false,              // Expired
						sqlmock.AnyArg(),   // ID - 自动生成
					).WillReturnResult(sqlmock.NewResult(1, 1))

				// 期望执行订单项插入
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `order_items`")).
					WithArgs(
						sqlmock.AnyArg(), // CreatedAt
						sqlmock.AnyArg(), // UpdatedAt
						nil,              // DeletedAt
						"order123",       // OrderID
						uint32(5001),     // ProductID
						int32(2),         // Quantity
						float32(99.99),   // Cost
					).WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
			expectedID:    "order123",
			expectedError: false,
		}}

	// 运行测试用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 设置模拟数据库
			db, mock := setupMockDB(t)
			tc.mockSetup(mock)

			// 创建实例并测试
			repo := NewOrderRepo(db)
			orderID, err := repo.CreateOrder(tc.order)
			log.Println(orderID)

			// 验证结果
			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expectedID, orderID)

			// 验证所有SQL预期已满足
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestListOrders(t *testing.T) {
	// 创建一个固定的时间用于测试
	fixedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	// 设置测试用例
	testCases := []struct {
		name           string
		userID         uint32
		mockSetup      func(sqlmock.Sqlmock)
		expectedOrders int
		expectedError  bool
	}{
		{
			name:   "user with multiple orders",
			userID: 1001,
			mockSetup: func(mock sqlmock.Sqlmock) {
				orderRows := sqlmock.NewRows([]string{
					"id", "created_at", "updated_at", "deleted_at",
					"order_id", "user_id", "user_currency", "email",
					"email", "street_address", "city", "state", "country", "zip_code",
					"paid", "expired",
				}).
					AddRow(1, fixedTime, fixedTime, nil,
						"order1", 1001, "USD", "test@example.com",
						"test@example.com", "123 Main St", "New York", "NY", "USA", 10001,
						false, false).
					AddRow(2, fixedTime, fixedTime, nil,
						"order2", 1001, "USD", "test@example.com",
						"test@example.com", "123 Main St", "New York", "NY", "USA", 10001,
						true, false)

				// 主查询 - 查找用户的所有订单
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `orders` WHERE user_id = ?")).
					WithArgs(1001).
					WillReturnRows(orderRows)

				// Preload查询 - 第一个订单的订单项
				allItemRows := sqlmock.NewRows([]string{
					"id", "created_at", "updated_at", "deleted_at",
					"order_id", "product_id", "quantity", "cost",
				}).
					// order1的订单项
					AddRow(1, fixedTime, fixedTime, nil,
						"order1", 5001, 2, 49.99).
					AddRow(2, fixedTime, fixedTime, nil,
						"order1", 5002, 1, 29.99).
					// order2的订单项
					AddRow(3, fixedTime, fixedTime, nil,
						"order2", 5003, 1, 199.99)

				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `order_items` WHERE `order_items`.`order_id` IN (?,?) AND `order_items`.`deleted_at` IS NULL ")).
					WithArgs("order1", "order2").
					WillReturnRows(allItemRows)
			},
			expectedOrders: 2,
			expectedError:  false,
		},
	}

	// 运行测试用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 设置模拟数据库

			db, mock := setupMockDB(t)
			db.Logger.LogMode(4)
			tc.mockSetup(mock)

			// 创建实例并测试
			repo := NewOrderRepo(db)
			orders, err := repo.ListOrders(tc.userID)

			// 验证结果
			if tc.expectedError {
				assert.Error(t, err)
				assert.Nil(t, orders)
			} else {
				assert.NoError(t, err)
				assert.Len(t, orders, tc.expectedOrders)
				// 验证第一个用例的订单项数量
				if tc.name == "user with multiple orders" && len(orders) > 1 {
					assert.Len(t, orders[0].OrderItems, 2) // 第一个订单有2个订单项
					assert.Len(t, orders[1].OrderItems, 1) // 第二个订单有1个订单项
				}
			}

			// 验证所有SQL预期已满足
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestMarkOrderPaid(t *testing.T) {
	// 设置测试用例
	testCases := []struct {
		name          string
		orderID       string
		mockSetup     func(sqlmock.Sqlmock)
		expectedError bool
	}{
		{
			name:    "successful update",
			orderID: "order123",
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("UPDATE `orders` SET `paid`=?,`updated_at`=? WHERE order_id = ? AND `orders`.`deleted_at` IS  NULL")).
					WithArgs(true, sqlmock.AnyArg(), "order123").
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
			expectedError: false,
		},
	}

	// 运行测试用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 设置模拟数据库
			db, mock := setupMockDB(t)
			tc.mockSetup(mock)

			// 创建实例并测试
			repo := NewOrderRepo(db)
			err := repo.MarkOrderPaid(tc.orderID)

			// 验证结果
			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// 验证所有SQL预期已满足
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
