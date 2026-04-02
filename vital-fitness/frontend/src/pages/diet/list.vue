<template>
	<view class="container">
		<!-- 顶部导航栏 -->
		<u-navbar title="饮食记录" fixed placeholder>
			<view slot="right">
				<u-icon name="plus" size="20" @click="addDiet"></u-icon>
			</view>
		</u-navbar>

		<!-- 今日饮食统计 -->
		<view class="diet-stats">
			<view class="stats-header">
				<text class="date">{{ today }}</text>
				<text class="calories-total">{{ totalCalories }} 卡路里</text>
			</view>

			<view class="meal-summary">
				<view class="meal-item" v-for="meal in mealSummary" :key="meal.type">
					<view class="meal-icon">
						<u-icon :name="meal.icon" :color="meal.color" size="24"></u-icon>
					</view>
					<view class="meal-info">
						<text class="meal-name">{{ meal.name }}</text>
						<text class="meal-calories">{{ meal.calories }}卡</text>
					</view>
					<view class="meal-count">
						<text class="count">{{ meal.count }}项</text>
					</view>
				</view>
			</view>

			<!-- 营养素分布 -->
			<view class="nutrition-summary">
				<view class="nutrition-item">
					<text class="nutrition-label">蛋白质</text>
					<text class="nutrition-value">{{ totalProtein }}g</text>
				</view>
				<view class="nutrition-item">
					<text class="nutrition-label">碳水</text>
					<text class="nutrition-value">{{ totalCarbs }}g</text>
				</view>
				<view class="nutrition-item">
					<text class="nutrition-label">脂肪</text>
					<text class="nutrition-value">{{ totalFat }}g</text>
				</view>
			</view>
		</view>

		<!-- 饮食记录列表 -->
		<view class="diet-list">
			<view class="list-header">
				<text class="header-title">详细记录</text>
			</view>

			<view v-for="(diet, index) in dietList" :key="index" class="diet-item">
				<view class="diet-header">
					<view class="meal-type">
						<u-icon :name="getMealIcon(diet.meal_type)" :color="getMealColor(diet.meal_type)" size="18"></u-icon>
						<text class="meal-name">{{ getMealName(diet.meal_type) }}</text>
					</view>
					<text class="diet-time">{{ formatTime(diet.created_at) }}</text>
				</view>

				<view class="diet-content">
					<view class="food-items">
						<view v-for="(food, foodIndex) in parseFoodItems(diet.food_items)" :key="foodIndex" class="food-item">
							<text class="food-name">{{ food.name }}</text>
							<text class="food-amount">{{ food.amount }}</text>
						</view>
					</view>

					<view class="diet-nutrition">
						<text class="calories">{{ diet.total_calories }}卡路里</text>
						<view class="macro-nutrients">
							<text v-if="diet.protein" class="nutrient">蛋白质: {{ diet.protein }}g</text>
							<text v-if="diet.carbs" class="nutrient">碳水: {{ diet.carbs }}g</text>
							<text v-if="diet.fat" class="nutrient">脂肪: {{ diet.fat }}g</text>
						</view>
					</view>

					<view v-if="diet.notes" class="diet-notes">
						<u-icon name="edit-pen" size="16" color="#999"></u-icon>
						<text class="notes-text">{{ diet.notes }}</text>
					</view>
				</view>
			</view>

			<!-- 空状态 -->
			<view v-if="dietList.length === 0" class="empty-state">
				<u-empty text="暂无饮食记录" mode="gift"></u-empty>
				<u-button type="primary" icon="plus" @click="addDiet" style="margin-top: 30rpx;">记录第一餐</u-button>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				today: '2026年3月13日',
				dietList: [
					{
						id: 1,
						meal_type: 'breakfast',
						food_items: '[{"name":"燕麦粥","amount":"1碗"},{"name":"鸡蛋","amount":"2个"},{"name":"牛奶","amount":"200ml"}]',
						total_calories: 420,
						protein: 25,
						carbs: 45,
						fat: 12,
						created_at: '2026-03-13 08:15:00',
						notes: '加入了一些坚果'
					},
					{
						id: 2,
						meal_type: 'lunch',
						food_items: '[{"name":"鸡胸肉","amount":"150g"},{"name":"糙米饭","amount":"1碗"},{"name":"西兰花","amount":"100g"}]',
						total_calories: 580,
						protein: 45,
						carbs: 35,
						fat: 15,
						created_at: '2026-03-13 12:30:00',
						notes: '少油烹饪'
					},
					{
						id: 3,
						meal_type: 'snack',
						food_items: '[{"name":"苹果","amount":"1个"},{"name":"杏仁","amount":"20颗"}]',
						total_calories: 180,
						protein: 5,
						carbs: 20,
						fat: 8,
						created_at: '2026-03-13 15:45:00',
						notes: ''
					}
				],
				mealSummary: [
					{ type: 'breakfast', name: '早餐', calories: 420, count: 3, icon: 'coffee', color: '#ff9900' },
					{ type: 'lunch', name: '午餐', calories: 580, count: 3, icon: 'restaurant', color: '#3c9cff' },
					{ type: 'dinner', name: '晚餐', calories: 0, count: 0, icon: 'moon', color: '#a78bfa' },
					{ type: 'snack', name: '加餐', calories: 180, count: 2, icon: 'gift', color: '#6bcf7f' }
				]
			}
		},
		computed: {
			totalCalories() {
				return this.dietList.reduce((sum, diet) => sum + diet.total_calories, 0)
			},
			totalProtein() {
				return this.dietList.reduce((sum, diet) => sum + (diet.protein || 0), 0)
			},
			totalCarbs() {
				return this.dietList.reduce((sum, diet) => sum + (diet.carbs || 0), 0)
			},
			totalFat() {
				return this.dietList.reduce((sum, diet) => sum + (diet.fat || 0), 0)
			}
		},
		methods: {
			addDiet() {
				uni.navigateTo({
					url: '/pages/diet/add'
				})
			},

			getMealName(mealType) {
				const mealNames = {
					'breakfast': '早餐',
					'lunch': '午餐',
					'dinner': '晚餐',
					'snack': '加餐'
				}
				return mealNames[mealType] || '未知'
			},

			getMealIcon(mealType) {
				const mealIcons = {
					'breakfast': 'coffee',
					'lunch': 'restaurant',
					'dinner': 'moon',
					'snack': 'gift'
				}
				return mealIcons[mealType] || 'food'
			},

			getMealColor(mealType) {
				const mealColors = {
					'breakfast': '#ff9900',
					'lunch': '#3c9cff',
					'dinner': '#a78bfa',
					'snack': '#6bcf7f'
				}
				return mealColors[mealType] || '#999'
			},

			parseFoodItems(foodItems) {
				try {
					return JSON.parse(foodItems)
				} catch (e) {
					return []
				}
			},

			formatTime(dateTimeString) {
				const date = new Date(dateTimeString)
				return `${date.getHours().toString().padStart(2, '0')}:${date.getMinutes().toString().padStart(2, '0')}`
			}
		}
	}
</script>

<style lang="scss" scoped>
.container {
	padding: 20rpx;
}

.diet-stats {
	background: white;
	border-radius: 20rpx;
	padding: 30rpx;
	margin-bottom: 20rpx;

	.stats-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 30rpx;

		.date {
			font-size: 32rpx;
			font-weight: bold;
			color: #333;
		}

		.calories-total {
			font-size: 28rpx;
			color: #ff9900;
			font-weight: bold;
		}
	}

	.meal-summary {
		.meal-item {
			display: flex;
			align-items: center;
			padding: 20rpx 0;
			border-bottom: 1rpx solid #eee;

			&:last-child {
				border-bottom: none;
			}

			.meal-icon {
				margin-right: 20rpx;
			}

			.meal-info {
				flex: 1;

				.meal-name {
					font-size: 28rpx;
					color: #333;
					display: block;
				}

				.meal-calories {
					font-size: 24rpx;
					color: #999;
				}
			}

			.meal-count {
				.count {
					font-size: 24rpx;
					color: #666;
				}
			}
		}
	}

	.nutrition-summary {
		display: flex;
		justify-content: space-around;
		margin-top: 30rpx;
		padding-top: 30rpx;
		border-top: 1rpx solid #eee;

		.nutrition-item {
			text-align: center;

			.nutrition-label {
				font-size: 24rpx;
				color: #999;
				display: block;
				margin-bottom: 10rpx;
			}

			.nutrition-value {
				font-size: 28rpx;
				font-weight: bold;
				color: #333;
			}
		}
	}
}

.diet-list {
	background: white;
	border-radius: 20rpx;
	overflow: hidden;

	.list-header {
		padding: 30rpx;
		border-bottom: 1rpx solid #eee;

		.header-title {
			font-size: 32rpx;
			font-weight: bold;
			color: #333;
		}
	}

	.diet-item {
		padding: 30rpx;
		border-bottom: 1rpx solid #eee;

		&:last-child {
			border-bottom: none;
		}

		.diet-header {
			display: flex;
			justify-content: space-between;
			align-items: center;
			margin-bottom: 20rpx;

			.meal-type {
				display: flex;
				align-items: center;

				.meal-name {
					font-size: 28rpx;
					font-weight: bold;
					color: #333;
					margin-left: 10rpx;
				}
			}

			.diet-time {
				font-size: 24rpx;
				color: #999;
			}
		}

		.diet-content {
			.food-items {
				margin-bottom: 20rpx;

				.food-item {
					display: flex;
					justify-content: space-between;
					padding: 10rpx 0;

					.food-name {
						font-size: 26rpx;
						color: #666;
					}

					.food-amount {
						font-size: 24rpx;
						color: #999;
					}
				}
			}

			.diet-nutrition {
				margin-bottom: 20rpx;

				.calories {
					font-size: 28rpx;
					font-weight: bold;
					color: #ff9900;
					display: block;
					margin-bottom: 10rpx;
				}

				.macro-nutrients {
					display: flex;
					flex-wrap: wrap;
					gap: 15rpx;

					.nutrient {
						font-size: 22rpx;
						color: #666;
						background: #f5f5f5;
						padding: 5rpx 15rpx;
						border-radius: 10rpx;
					}
				}
			}

			.diet-notes {
				display: flex;
				align-items: flex-start;

				.notes-text {
					font-size: 24rpx;
					color: #999;
					margin-left: 10rpx;
					line-height: 1.4;
				}
			}
		}
	}

	.empty-state {
		text-align: center;
		padding: 100rpx 0;
	}
}
</style>