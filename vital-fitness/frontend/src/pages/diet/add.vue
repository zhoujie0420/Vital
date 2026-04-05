<template>
	<view class="container">
		<u-navbar title="记录饮食" fixed placeholder>
			<view slot="left">
				<u-icon name="arrow-left" size="20" @click="goBack"></u-icon>
			</view>
		</u-navbar>

		<view class="form-container">
			<!-- 餐次选择 -->
			<view class="form-item">
				<text class="label">餐次</text>
				<view class="meal-tabs">
					<text v-for="m in meals" :key="m.type" class="meal-tab"
						:class="{ active: mealType === m.type }" @tap="mealType = m.type">
						{{ m.name }}
					</text>
				</view>
			</view>

			<!-- 搜索食物 -->
			<view class="form-item">
				<text class="label">选择食物</text>
				<input class="search-input" v-model="keyword" placeholder="搜索食物..." @input="searchFoods" />

				<!-- 分类 -->
				<view class="cat-tabs">
					<text v-for="c in categories" :key="c" class="cat-tab"
						:class="{ active: currentCat === c }" @tap="currentCat = c; searchFoods()">
						{{ c }}
					</text>
				</view>

				<!-- 食物列表 -->
				<view class="food-list">
					<view v-for="f in foodList" :key="f.id" class="food-item" @tap="selectFood(f)">
						<view class="food-info">
							<text class="food-name">{{ f.name }}</text>
							<text class="food-detail">{{ f.calories }}kcal | 蛋白{{ f.protein }}g | 碳水{{ f.carbs }}g | 脂肪{{ f.fat }}g</text>
						</view>
						<text class="food-unit">/ {{ f.unit }}</text>
					</view>
				</view>

				<view class="add-custom" @tap="showCustom = true">
					<text>+ 添加自定义食物</text>
				</view>
			</view>

			<!-- 已选食物 -->
			<view class="form-item" v-if="selectedFoods.length > 0">
				<text class="label">已选食物</text>
				<view v-for="(sf, i) in selectedFoods" :key="i" class="selected-food">
					<text class="sf-name">{{ sf.name }}</text>
					<text class="sf-cal">{{ sf.calories }}kcal</text>
					<text class="sf-remove" @tap="selectedFoods.splice(i, 1)">x</text>
				</view>
				<view class="total-row">
					<text>合计: {{ totalCalories }}kcal | 蛋白{{ totalProtein }}g | 碳水{{ totalCarbs }}g | 脂肪{{ totalFat }}g</text>
				</view>
			</view>

			<!-- 备注 -->
			<view class="form-item" v-if="selectedFoods.length > 0">
				<textarea v-model="notes" placeholder="备注（可选）" class="notes-input" />
			</view>

			<button v-if="selectedFoods.length > 0" class="save-btn" type="primary" :loading="saving" @tap="save">
				{{ saving ? '保存中...' : '保存记录' }}
			</button>
		</view>

		<!-- 自定义食物弹窗 -->
		<u-popup :show="showCustom" mode="bottom" @close="showCustom = false">
			<view class="custom-popup">
				<text class="popup-title">添加自定义食物</text>
				<input v-model="customFood.name" placeholder="食物名称" class="popup-input" />
				<input v-model="customFood.calories" type="digit" placeholder="热量(kcal/100g)" class="popup-input" />
				<input v-model="customFood.protein" type="digit" placeholder="蛋白质(g/100g)" class="popup-input" />
				<input v-model="customFood.carbs" type="digit" placeholder="碳水(g/100g)" class="popup-input" />
				<input v-model="customFood.fat" type="digit" placeholder="脂肪(g/100g)" class="popup-input" />
				<button type="primary" @tap="saveCustomFood">保存食物</button>
			</view>
		</u-popup>
	</view>
</template>

<script>
	import { getFoods, createFood, createDietRecord } from '../../api/diet'

	export default {
		data() {
			return {
				saving: false, showCustom: false,
				mealType: 'lunch', keyword: '', currentCat: '全部',
				meals: [
					{ type: 'breakfast', name: '早餐' },
					{ type: 'lunch', name: '午餐' },
					{ type: 'dinner', name: '晚餐' },
					{ type: 'snack', name: '加餐' }
				],
				categories: ['全部', '主食', '肉类', '蔬菜', '水果', '奶制品', '零食'],
				foodList: [],
				selectedFoods: [],
				notes: '',
				customFood: { name: '', calories: '', protein: '', carbs: '', fat: '' }
			}
		},
		computed: {
			totalCalories() { return this.selectedFoods.reduce((s, f) => s + f.calories, 0).toFixed(0) },
			totalProtein() { return this.selectedFoods.reduce((s, f) => s + f.protein, 0).toFixed(1) },
			totalCarbs() { return this.selectedFoods.reduce((s, f) => s + f.carbs, 0).toFixed(1) },
			totalFat() { return this.selectedFoods.reduce((s, f) => s + f.fat, 0).toFixed(1) }
		},
		onLoad() { this.searchFoods() },
		methods: {
			goBack() { uni.navigateBack() },
			async searchFoods() {
				try {
					const res = await getFoods({
						keyword: this.keyword,
						category: this.currentCat === '全部' ? '' : this.currentCat
					})
					this.foodList = res.data || []
				} catch (e) {}
			},
			selectFood(f) {
				this.selectedFoods.push({ ...f })
			},
			async saveCustomFood() {
				if (!this.customFood.name || !this.customFood.calories) {
					uni.showToast({ title: '请填写名称和热量', icon: 'none' })
					return
				}
				try {
					const res = await createFood({
						name: this.customFood.name,
						calories: parseFloat(this.customFood.calories),
						protein: parseFloat(this.customFood.protein) || 0,
						carbs: parseFloat(this.customFood.carbs) || 0,
						fat: parseFloat(this.customFood.fat) || 0
					})
					this.foodList.unshift(res.data)
					this.showCustom = false
					this.customFood = { name: '', calories: '', protein: '', carbs: '', fat: '' }
					uni.showToast({ title: '添加成功', icon: 'success' })
				} catch (e) {}
			},
			async save() {
				this.saving = true
				try {
					await createDietRecord({
						record_date: new Date().toISOString(),
						meal_type: this.mealType,
						food_items: JSON.stringify(this.selectedFoods.map(f => ({ name: f.name, calories: f.calories }))),
						total_calories: parseFloat(this.totalCalories),
						protein: parseFloat(this.totalProtein),
						carbs: parseFloat(this.totalCarbs),
						fat: parseFloat(this.totalFat),
						notes: this.notes
					})
					uni.showToast({ title: '保存成功', icon: 'success' })
					setTimeout(() => uni.navigateBack(), 1000)
				} catch (e) {} finally { this.saving = false }
			}
		}
	}
</script>

<style lang="scss" scoped>
.container { padding: 20rpx; }
.form-container {
	.form-item {
		background: white; border-radius: 20rpx; padding: 30rpx; margin-bottom: 20rpx;
		.label { font-size: 32rpx; font-weight: bold; color: #333; margin-bottom: 20rpx; display: block; }
		.meal-tabs, .cat-tabs { display: flex; flex-wrap: wrap; gap: 15rpx; margin-bottom: 15rpx; }
		.meal-tab, .cat-tab {
			padding: 10rpx 24rpx; border-radius: 30rpx; background: #f5f5f5; font-size: 24rpx; color: #666;
			&.active { background: #3c9cff; color: white; }
		}
		.search-input { border: 1rpx solid #eee; border-radius: 10rpx; padding: 15rpx; font-size: 28rpx; margin-bottom: 15rpx; }
		.food-list { max-height: 400rpx; overflow-y: auto;
			.food-item {
				display: flex; justify-content: space-between; align-items: center;
				padding: 20rpx 0; border-bottom: 1rpx solid #f5f5f5;
				.food-info {
					.food-name { font-size: 28rpx; color: #333; display: block; }
					.food-detail { font-size: 20rpx; color: #999; }
				}
				.food-unit { font-size: 22rpx; color: #999; }
			}
		}
		.add-custom { text-align: center; padding: 20rpx; color: #3c9cff; font-size: 26rpx; }
		.selected-food {
			display: flex; align-items: center; padding: 15rpx 0; border-bottom: 1rpx solid #f5f5f5;
			.sf-name { flex: 1; font-size: 28rpx; color: #333; }
			.sf-cal { font-size: 24rpx; color: #ff9900; margin-right: 20rpx; }
			.sf-remove { color: #ff4d4f; font-size: 28rpx; padding: 0 10rpx; }
		}
		.total-row { margin-top: 15rpx; padding-top: 15rpx; border-top: 1rpx solid #eee; font-size: 24rpx; color: #3c9cff; }
		.notes-input { width: 100%; height: 100rpx; border: 1rpx solid #eee; border-radius: 10rpx; padding: 15rpx; font-size: 26rpx; }
	}
	.save-btn { margin-top: 20rpx; border-radius: 15rpx; font-size: 32rpx; }
}
.custom-popup {
	padding: 40rpx;
	.popup-title { font-size: 32rpx; font-weight: bold; text-align: center; margin-bottom: 30rpx; display: block; }
	.popup-input { border: 1rpx solid #eee; border-radius: 10rpx; padding: 15rpx; font-size: 28rpx; margin-bottom: 20rpx; }
}
</style>
