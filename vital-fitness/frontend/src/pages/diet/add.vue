<template>
	<view class="page">
		<view class="nav-bar">
			<text class="nav-back" @tap="goBack">‹ 返回</text>
			<text class="nav-title">记录饮食</text>
			<text class="nav-placeholder"></text>
		</view>

		<!-- 餐次选择 -->
		<view class="card">
			<text class="card-label">餐次</text>
			<view class="segment">
				<text v-for="m in meals" :key="m.type" class="segment-item"
					:class="{ active: mealType === m.type }" @tap="mealType = m.type">
					{{ m.name }}
				</text>
			</view>
		</view>

		<!-- 搜索食物 -->
		<view class="card">
			<text class="card-label">选择食物</text>
			<view class="search-bar">
				<text class="search-icon">🔍</text>
				<input class="search-input" v-model="keyword" placeholder="搜索食物" @input="searchFoods" />
			</view>

			<scroll-view scroll-x class="cat-scroll">
				<text v-for="c in categories" :key="c" class="cat-pill"
					:class="{ active: currentCat === c }" @tap="currentCat = c; searchFoods()">
					{{ c }}
				</text>
			</scroll-view>

			<view class="food-list">
				<view v-for="f in foodList" :key="f.id" class="food-row" @tap="selectFood(f)">
					<view class="food-info">
						<text class="food-name">{{ f.name }}</text>
						<text class="food-detail">{{ f.calories }}kcal · 蛋白{{ f.protein }}g · 碳水{{ f.carbs }}g · 脂肪{{ f.fat }}g</text>
					</view>
					<text class="food-add">+</text>
				</view>
			</view>

			<view class="custom-link" @tap="showCustom = true">
				<text>+ 自定义食物</text>
			</view>
		</view>

		<!-- 已选食物 -->
		<view class="card" v-if="selectedFoods.length > 0">
			<text class="card-label">已选食物</text>
			<view v-for="(sf, i) in selectedFoods" :key="i" class="selected-row">
				<text class="sf-name">{{ sf.name }}</text>
				<text class="sf-cal">{{ sf.calories }}kcal</text>
				<text class="sf-remove" @tap="selectedFoods.splice(i, 1)">✕</text>
			</view>
			<view class="total-bar">
				<text class="total-text">合计 {{ totalCalories }}千卡</text>
				<text class="total-detail">蛋白{{ totalProtein }}g · 碳水{{ totalCarbs }}g · 脂肪{{ totalFat }}g</text>
			</view>
		</view>

		<!-- 备注 -->
		<view class="card" v-if="selectedFoods.length > 0">
			<textarea v-model="notes" placeholder="备注（可选）" class="notes-area" />
		</view>

		<view class="bottom-action" v-if="selectedFoods.length > 0">
			<view class="save-btn" :class="{ loading: saving }" @tap="save">
				<text>{{ saving ? '保存中...' : '保存记录' }}</text>
			</view>
		</view>

		<!-- 自定义食物弹窗 -->
		<u-popup :show="showCustom" mode="bottom" round="24" @close="showCustom = false">
			<view class="popup">
				<text class="popup-title">自定义食物</text>
				<view class="popup-field">
					<text class="popup-label">名称</text>
					<input v-model="customFood.name" placeholder="食物名称" class="popup-input" />
				</view>
				<view class="popup-field">
					<text class="popup-label">热量</text>
					<input v-model="customFood.calories" type="digit" placeholder="kcal / 100g" class="popup-input" />
				</view>
				<view class="popup-field">
					<text class="popup-label">蛋白质</text>
					<input v-model="customFood.protein" type="digit" placeholder="g / 100g" class="popup-input" />
				</view>
				<view class="popup-field">
					<text class="popup-label">碳水</text>
					<input v-model="customFood.carbs" type="digit" placeholder="g / 100g" class="popup-input" />
				</view>
				<view class="popup-field">
					<text class="popup-label">脂肪</text>
					<input v-model="customFood.fat" type="digit" placeholder="g / 100g" class="popup-input" />
				</view>
				<view class="popup-btn" @tap="saveCustomFood">
					<text>保存</text>
				</view>
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
			selectFood(f) { this.selectedFoods.push({ ...f }) },
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
.page {
	padding: 0 32rpx;
	padding-top: 120rpx;
	padding-bottom: 160rpx;
	min-height: 100vh;
	background: #f2f2f7;
}

.nav-bar {
	display: flex;
	align-items: center;
	justify-content: space-between;
	margin-bottom: 28rpx;

	.nav-back {
		font-size: 32rpx;
		color: #007aff;
		font-weight: 500;
	}
	.nav-title {
		font-size: 34rpx;
		font-weight: 600;
		color: #1c1c1e;
	}
	.nav-placeholder { width: 80rpx; }
}

.card {
	background: #fff;
	border-radius: 20rpx;
	padding: 28rpx 32rpx;
	margin-bottom: 16rpx;
	box-shadow: 0 2rpx 16rpx rgba(0, 0, 0, 0.04);

	.card-label {
		display: block;
		font-size: 26rpx;
		font-weight: 600;
		color: #8e8e93;
		text-transform: uppercase;
		letter-spacing: 1rpx;
		margin-bottom: 20rpx;
	}
}

.segment {
	display: flex;
	background: #f2f2f7;
	border-radius: 16rpx;
	padding: 4rpx;

	.segment-item {
		flex: 1;
		text-align: center;
		padding: 16rpx 0;
		font-size: 28rpx;
		font-weight: 500;
		color: #636366;
		border-radius: 14rpx;
		transition: all 0.2s;

		&.active {
			background: #fff;
			color: #1c1c1e;
			font-weight: 600;
			box-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.08);
		}
	}
}

.search-bar {
	display: flex;
	align-items: center;
	background: #f2f2f7;
	border-radius: 16rpx;
	padding: 16rpx 20rpx;
	margin-bottom: 16rpx;

	.search-icon { font-size: 28rpx; margin-right: 12rpx; }
	.search-input { flex: 1; font-size: 28rpx; color: #1c1c1e; }
}

.cat-scroll {
	white-space: nowrap;
	margin-bottom: 16rpx;

	.cat-pill {
		display: inline-block;
		padding: 10rpx 24rpx;
		border-radius: 20rpx;
		background: #f2f2f7;
		font-size: 26rpx;
		color: #636366;
		margin-right: 12rpx;
		font-weight: 500;

		&.active {
			background: #007aff;
			color: #fff;
		}
	}
}

.food-list {
	max-height: 400rpx;
	overflow-y: auto;

	.food-row {
		display: flex;
		align-items: center;
		padding: 20rpx 0;
		border-bottom: 1rpx solid #f2f2f7;

		&:last-child { border-bottom: none; }
		&:active { opacity: 0.6; }

		.food-info {
			flex: 1;
			.food-name { display: block; font-size: 30rpx; font-weight: 500; color: #1c1c1e; }
			.food-detail { display: block; font-size: 22rpx; color: #8e8e93; margin-top: 4rpx; }
		}
		.food-add {
			width: 56rpx;
			height: 56rpx;
			line-height: 52rpx;
			text-align: center;
			background: #f2f2f7;
			border-radius: 50%;
			font-size: 36rpx;
			color: #007aff;
			font-weight: 300;
		}
	}
}

.custom-link {
	text-align: center;
	padding: 20rpx 0 4rpx;
	font-size: 28rpx;
	color: #007aff;
	font-weight: 500;
}

.selected-row {
	display: flex;
	align-items: center;
	padding: 16rpx 0;
	border-bottom: 1rpx solid #f2f2f7;

	.sf-name { flex: 1; font-size: 30rpx; color: #1c1c1e; font-weight: 500; }
	.sf-cal { font-size: 26rpx; color: #8e8e93; margin-right: 20rpx; }
	.sf-remove { font-size: 24rpx; color: #c7c7cc; padding: 8rpx; }
}

.total-bar {
	margin-top: 20rpx;
	padding-top: 20rpx;
	border-top: 1rpx solid #f2f2f7;

	.total-text {
		display: block;
		font-size: 34rpx;
		font-weight: 700;
		color: #1c1c1e;
	}
	.total-detail {
		display: block;
		font-size: 24rpx;
		color: #8e8e93;
		margin-top: 4rpx;
	}
}

.notes-area {
	width: 100%;
	height: 120rpx;
	font-size: 28rpx;
	color: #1c1c1e;
	line-height: 1.6;
}

.bottom-action {
	position: fixed;
	bottom: 0;
	left: 0;
	right: 0;
	padding: 20rpx 32rpx;
	padding-bottom: calc(20rpx + env(safe-area-inset-bottom));
	background: rgba(242, 242, 247, 0.9);
	backdrop-filter: blur(20px);

	.save-btn {
		background: #007aff;
		color: #fff;
		text-align: center;
		padding: 28rpx 0;
		border-radius: 16rpx;
		font-size: 32rpx;
		font-weight: 600;

		&:active { opacity: 0.8; }
		&.loading { opacity: 0.6; }
	}
}

.popup {
	padding: 40rpx 32rpx;
	padding-bottom: calc(40rpx + env(safe-area-inset-bottom));

	.popup-title {
		display: block;
		font-size: 36rpx;
		font-weight: 700;
		color: #1c1c1e;
		text-align: center;
		margin-bottom: 32rpx;
	}
	.popup-field {
		margin-bottom: 20rpx;

		.popup-label {
			display: block;
			font-size: 26rpx;
			color: #8e8e93;
			font-weight: 500;
			margin-bottom: 8rpx;
		}
		.popup-input {
			background: #f2f2f7;
			border-radius: 12rpx;
			padding: 20rpx;
			font-size: 30rpx;
			color: #1c1c1e;
		}
	}
	.popup-btn {
		background: #007aff;
		color: #fff;
		text-align: center;
		padding: 28rpx 0;
		border-radius: 16rpx;
		font-size: 32rpx;
		font-weight: 600;
		margin-top: 12rpx;

		&:active { opacity: 0.8; }
	}
}
</style>
