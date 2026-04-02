<template>
	<view class="container">
		<!-- 顶部导航栏 -->
		<u-navbar title="记录饮食" fixed placeholder>
			<view slot="left">
				<u-icon name="arrow-left" size="20" @click="goBack"></u-icon>
			</view>
			<view slot="right">
				<text class="nav-btn" @click="saveDiet">保存</text>
			</view>
		</u-navbar>

		<view class="form-container">
			<!-- 餐次选择 -->
			<view class="form-item">
				<text class="label">餐次</text>
				<view class="meal-types">
					<view
						v-for="meal in mealTypes"
						:key="meal.type"
						class="meal-item"
						:class="{ active: dietForm.meal_type === meal.type }"
						@click="selectMealType(meal.type)"
					>
						<u-icon :name="meal.icon" :color="dietForm.meal_type === meal.type ? '#fff' : meal.color" size="20"></u-icon>
						<text class="meal-name">{{ meal.name }}</text>
					</view>
				</view>
			</view>

			<!-- 记录时间 -->
			<view class="form-item">
				<text class="label">记录时间</text>
				<u-datetime-picker
					:show="showTimePicker"
					v-model="dietForm.record_time"
					mode="datetime"
					@confirm="confirmTime"
					@cancel="showTimePicker = false"
				></u-datetime-picker>
				<u-cell
					title="记录时间"
					:value="formatDateTime(dietForm.record_time)"
					isLink
					@click="showTimePicker = true"
				></u-cell>
			</view>

			<!-- 食物列表 -->
			<view class="form-item">
				<view class="section-header">
					<text class="label">食物列表</text>
					<u-icon name="plus" size="20" color="#3c9cff" @click="addFoodItem"></u-icon>
				</view>

				<view v-for="(food, index) in foodItems" :key="index" class="food-item">
					<u-icon
						name="close"
						size="16"
						color="#ff6b6b"
						class="remove-btn"
						@click="removeFoodItem(index)"
					></u-icon>

					<u-input
						v-model="food.name"
						placeholder="食物名称"
						border="none"
						style="flex: 1; margin-right: 20rpx;"
					/>

					<u-input
						v-model="food.amount"
						placeholder="份量"
						border="none"
						style="width: 120rpx; margin-right: 20rpx;"
					/>

					<u-button
						type="primary"
						size="mini"
						@click="searchFoodInfo(index)"
					>查询</u-button>
				</view>

				<view v-if="foodItems.length === 0" class="empty-food-list">
					<text>点击上方+号添加食物</text>
				</view>
			</view>

			<!-- 营养信息（自动计算） -->
			<view class="form-item">
				<text class="label">营养信息</text>

				<view class="nutrition-grid">
					<view class="nutrition-item">
						<text class="nutrition-label">总卡路里</text>
						<view class="nutrition-value">
							<u-input
								v-model="dietForm.total_calories"
								type="digit"
								placeholder="0"
								border="none"
								style="text-align: center;"
							/>
							<text class="unit">kcal</text>
						</view>
					</view>

					<view class="nutrition-item">
						<text class="nutrition-label">蛋白质</text>
						<view class="nutrition-value">
							<u-input
								v-model="dietForm.protein"
								type="digit"
								placeholder="0"
								border="none"
								style="text-align: center;"
							/>
							<text class="unit">g</text>
						</view>
					</view>

					<view class="nutrition-item">
						<text class="nutrition-label">碳水化合物</text>
						<view class="nutrition-value">
							<u-input
								v-model="dietForm.carbs"
								type="digit"
								placeholder="0"
								border="none"
								style="text-align: center;"
							/>
							<text class="unit">g</text>
						</view>
					</view>

					<view class="nutrition-item">
						<text class="nutrition-label">脂肪</text>
						<view class="nutrition-value">
							<u-input
								v-model="dietForm.fat"
								type="digit"
								placeholder="0"
								border="none"
								style="text-align: center;"
							/>
							<text class="unit">g</text>
						</view>
					</view>
				</view>
			</view>

			<!-- 饮水量 -->
			<view class="form-item">
				<text class="label">饮水量</text>
				<view class="water-intake">
					<text class="intake-label">今日饮水</text>
					<u-number-box
						v-model="dietForm.water_intake"
						:min="0"
						:step="100"
						label="ml"
					></u-number-box>
				</view>
			</view>

			<!-- 备注 -->
			<view class="form-item">
				<text class="label">备注</text>
				<u-textarea
					v-model="dietForm.notes"
					placeholder="记录烹饪方式、特殊说明等..."
					:height="120"
				/>
			</view>
		</view>

		<!-- 食物数据库搜索弹窗 -->
		<u-popup :show="showFoodSearch" mode="bottom" @close="showFoodSearch = false">
			<view class="food-search-popup">
				<view class="popup-header">
					<text class="popup-title">搜索食物</text>
					<u-icon name="close" size="20" @click="showFoodSearch = false"></u-icon>
				</view>

				<view class="search-box">
					<u-search
						v-model="searchKeyword"
						placeholder="搜索食物名称"
						@search="searchFoods"
						@custom="searchFoods"
					></u-search>
				</view>

				<view class="search-results">
					<view
						v-for="food in searchResults"
						:key="food.id"
						class="result-item"
						@click="selectFood(food)"
					>
						<view class="food-info">
							<text class="food-name">{{ food.name }}</text>
							<text class="food-calories">{{ food.calories }}kcal/100g</text>
						</view>
						<u-icon name="checkmark" color="#3c9cff" v-if="selectedFoodId === food.id"></u-icon>
					</view>

					<view v-if="searchResults.length === 0 && searchKeyword" class="no-results">
						<u-empty text="未找到相关食物" mode="search"></u-empty>
					</view>

					<view v-if="!searchKeyword" class="search-tip">
						<u-empty text="输入关键词搜索食物" mode="search"></u-empty>
					</view>
				</view>
			</view>
		</u-popup>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				showTimePicker: false,
				showFoodSearch: false,
				searchKeyword: '',
				selectedFoodId: null,
				currentFoodIndex: 0,

				dietForm: {
					meal_type: 'breakfast',
					record_time: Date.now(),
					total_calories: '',
					protein: '',
					carbs: '',
					fat: '',
					water_intake: 0,
					notes: ''
				},

				foodItems: [
					{ name: '', amount: '' }
				],

				mealTypes: [
					{ type: 'breakfast', name: '早餐', icon: 'coffee', color: '#ff9900' },
					{ type: 'lunch', name: '午餐', icon: 'restaurant', color: '#3c9cff' },
					{ type: 'dinner', name: '晚餐', icon: 'moon', color: '#a78bfa' },
					{ type: 'snack', name: '加餐', icon: 'gift', color: '#6bcf7f' }
				],

				searchResults: [
					{ id: 1, name: '鸡胸肉', calories: 165, protein: 31, carbs: 0, fat: 3.6 },
					{ id: 2, name: '三文鱼', calories: 208, protein: 20, carbs: 0, fat: 13 },
					{ id: 3, name: '牛肉', calories: 250, protein: 26, carbs: 0, fat: 15 },
					{ id: 4, name: '鸡蛋', calories: 155, protein: 13, carbs: 1.1, fat: 11 },
					{ id: 5, name: '燕麦', calories: 389, protein: 16.9, carbs: 66.3, fat: 6.9 }
				]
			}
		},
		methods: {
			goBack() {
				uni.navigateBack()
			},

			saveDiet() {
				if (this.foodItems.every(item => !item.name.trim())) {
					uni.showToast({
						title: '请至少添加一种食物',
						icon: 'none'
					})
					return
				}

				// 这里应该调用API保存饮食记录
				uni.showToast({
					title: '保存成功',
					icon: 'success'
				})

				// 返回上一页
				setTimeout(() => {
					uni.navigateBack()
				}, 1000)
			},

			selectMealType(type) {
				this.dietForm.meal_type = type
			},

			confirmTime(e) {
				this.dietForm.record_time = e.value
				this.showTimePicker = false
			},

			addFoodItem() {
				this.foodItems.push({ name: '', amount: '' })
			},

			removeFoodItem(index) {
				if (this.foodItems.length > 1) {
					this.foodItems.splice(index, 1)
				} else {
					this.foodItems[0] = { name: '', amount: '' }
				}
			},

			searchFoodInfo(index) {
				this.currentFoodIndex = index
				this.showFoodSearch = true
				this.searchKeyword = this.foodItems[index].name
				if (this.searchKeyword) {
					this.searchFoods()
				}
			},

			searchFoods() {
				// 实际项目中这里应该调用食物数据库API
				// 这里模拟搜索结果
				if (this.searchKeyword) {
					this.searchResults = this.searchResults.filter(food =>
						food.name.includes(this.searchKeyword)
					)
				} else {
					this.searchResults = [
						{ id: 1, name: '鸡胸肉', calories: 165, protein: 31, carbs: 0, fat: 3.6 },
						{ id: 2, name: '三文鱼', calories: 208, protein: 20, carbs: 0, fat: 13 },
						{ id: 3, name: '牛肉', calories: 250, protein: 26, carbs: 0, fat: 15 },
						{ id: 4, name: '鸡蛋', calories: 155, protein: 13, carbs: 1.1, fat: 11 },
						{ id: 5, name: '燕麦', calories: 389, protein: 16.9, carbs: 66.3, fat: 6.9 }
					]
				}
			},

			selectFood(food) {
				this.selectedFoodId = food.id
				this.foodItems[this.currentFoodIndex].name = food.name

				// 自动填充营养信息（简化处理）
				this.dietForm.total_calories = food.calories
				this.dietForm.protein = food.protein
				this.dietForm.carbs = food.carbs
				this.dietForm.fat = food.fat

				// 1秒后关闭弹窗
				setTimeout(() => {
					this.showFoodSearch = false
				}, 500)
			},

			formatDateTime(timestamp) {
				const date = new Date(timestamp)
				return `${date.getFullYear()}-${(date.getMonth() + 1).toString().padStart(2, '0')}-${date.getDate().toString().padStart(2, '0')} ${date.getHours().toString().padStart(2, '0')}:${date.getMinutes().toString().padStart(2, '0')}`
			}
		}
	}
</script>

<style lang="scss" scoped>
.container {
	padding: 20rpx;
}

.nav-btn {
	color: #3c9cff;
	font-size: 28rpx;
}

.form-container {
	.form-item {
		background: white;
		border-radius: 20rpx;
		padding: 0 30rpx 30rpx;
		margin-bottom: 20rpx;

		.section-header {
			display: flex;
			justify-content: space-between;
			align-items: center;
			padding: 30rpx 0 20rpx;

			.label {
				font-size: 32rpx;
				font-weight: bold;
				color: #333;
			}
		}

		.label {
			display: block;
			font-size: 32rpx;
			font-weight: bold;
			color: #333;
			padding: 30rpx 0 20rpx;
		}

		.meal-types {
			display: flex;
			justify-content: space-between;
			flex-wrap: wrap;
			gap: 20rpx;

			.meal-item {
				flex: 1;
				min-width: 25%;
				display: flex;
				flex-direction: column;
				align-items: center;
				justify-content: center;
				padding: 20rpx;
				border-radius: 10rpx;
				background: #f5f5f5;
				color: #666;

				&.active {
					background: #3c9cff;
					color: white;
				}

				.meal-name {
					font-size: 24rpx;
					margin-top: 10rpx;
				}
			}
		}

		.food-item {
			display: flex;
			align-items: center;
			padding: 20rpx 0;
			border-bottom: 1rpx dashed #eee;
			position: relative;

			&:last-child {
				border-bottom: none;
			}

			.remove-btn {
				position: absolute;
				top: 10rpx;
				right: 0;
			}
		}

		.empty-food-list {
			text-align: center;
			padding: 40rpx 0;
			color: #999;
		}

		.nutrition-grid {
			display: grid;
			grid-template-columns: repeat(2, 1fr);
			gap: 20rpx;

			.nutrition-item {
				background: #f5f5f5;
				border-radius: 10rpx;
				padding: 20rpx;
				text-align: center;

				.nutrition-label {
					font-size: 24rpx;
					color: #666;
					display: block;
					margin-bottom: 10rpx;
				}

				.nutrition-value {
					display: flex;
					align-items: center;
					justify-content: center;

					.unit {
						font-size: 24rpx;
						color: #999;
						margin-left: 5rpx;
					}
				}
			}
		}

		.water-intake {
			display: flex;
			align-items: center;
			justify-content: space-between;

			.intake-label {
				font-size: 28rpx;
				color: #333;
			}
		}
	}
}

.food-search-popup {
	.popup-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 30rpx;
		border-bottom: 1rpx solid #eee;

		.popup-title {
			font-size: 32rpx;
			font-weight: bold;
		}
	}

	.search-box {
		padding: 20rpx 30rpx;
	}

	.search-results {
		max-height: 500rpx;
		overflow-y: auto;

		.result-item {
			display: flex;
			justify-content: space-between;
			align-items: center;
			padding: 30rpx;
			border-bottom: 1rpx solid #eee;

			.food-info {
				.food-name {
					font-size: 28rpx;
					color: #333;
					display: block;
					margin-bottom: 10rpx;
				}

				.food-calories {
					font-size: 24rpx;
					color: #999;
				}
			}
		}

		.no-results,
		.search-tip {
			padding: 50rpx 0;
		}
	}
}
</style>