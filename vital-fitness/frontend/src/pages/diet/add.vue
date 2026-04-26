<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="nav-bar">
			<text class="nav-back" @tap="goBack">‹ 返回</text>
			<text class="nav-title">记录饮食</text>
			<view class="nav-placeholder"></view>
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

			<scroll-view scroll-x class="cat-scroll" :show-scrollbar="false">
				<view class="cat-bar">
					<text v-for="c in categories" :key="c" class="cat-pill"
						:class="{ active: currentCat === c }" @tap="currentCat = c; searchFoods()">
						{{ c }}
					</text>
				</view>
			</scroll-view>

			<view class="food-list">
				<view v-for="f in foodList" :key="f.id" class="food-row" @tap="selectFood(f)">
					<view class="food-info">
						<text class="food-name">{{ f.name }}</text>
						<text class="food-detail">{{ f.calories }}kcal · 蛋白{{ f.protein }}g · 碳水{{ f.carbs }}g · 脂肪{{ f.fat }}g</text>
					</view>
					<view class="food-add-btn">
						<text class="food-add-icon">+</text>
					</view>
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
				<view class="sf-remove" @tap="selectedFoods.splice(i, 1)">
					<text class="sf-remove-icon">✕</text>
				</view>
			</view>
			<view class="total-bar">
				<text class="total-val">{{ totalCalories }} 千卡</text>
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
		<view class="mask" v-if="showCustom" @tap="showCustom = false"></view>
		<view class="custom-sheet" v-if="showCustom">
			<view class="sheet-handle"></view>
			<text class="sheet-title">自定义食物</text>
			<view class="sheet-field">
				<text class="sheet-label">名称</text>
				<view class="sheet-input-wrap">
					<input v-model="customFood.name" placeholder="食物名称" class="sheet-input" />
				</view>
			</view>
			<view class="sheet-field">
				<text class="sheet-label">热量</text>
				<view class="sheet-input-wrap">
					<input v-model="customFood.calories" type="digit" placeholder="kcal / 100g" class="sheet-input" />
				</view>
			</view>
			<view class="sheet-field">
				<text class="sheet-label">蛋白质</text>
				<view class="sheet-input-wrap">
					<input v-model="customFood.protein" type="digit" placeholder="g / 100g" class="sheet-input" />
				</view>
			</view>
			<view class="sheet-field">
				<text class="sheet-label">碳水</text>
				<view class="sheet-input-wrap">
					<input v-model="customFood.carbs" type="digit" placeholder="g / 100g" class="sheet-input" />
				</view>
			</view>
			<view class="sheet-field">
				<text class="sheet-label">脂肪</text>
				<view class="sheet-input-wrap">
					<input v-model="customFood.fat" type="digit" placeholder="g / 100g" class="sheet-input" />
				</view>
			</view>
			<view class="sheet-save" @tap="saveCustomFood">
				<text>保存</text>
			</view>
		</view>
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
			topPadding() {
				const app = getApp()
				return (app.globalData?.customBarHeight || 88) + 8
			},
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
@import '../../styles/variables.scss';

.page {
	padding: 0 $spacing-xl;
	padding-bottom: 160rpx;
	min-height: 100vh;
	background: $color-bg;
}

.nav-bar {
	@include nav-bar;
	.nav-back { font-size: $font-callout; color: $color-primary; font-weight: 500; }
	.nav-title { font-size: $font-headline; font-weight: 600; color: $color-label; }
	.nav-placeholder { width: 80rpx; }
}

.card {
	@include card;
	margin-bottom: $spacing-md;
	.card-label { @include card-label; }
}

// --- Segment ---
.segment { @include segment-control; }

// --- Search ---
.search-bar {
	@include input-field;
	margin-bottom: $spacing-md;

	.search-icon { font-size: 28rpx; margin-right: $spacing-sm; }
	.search-input { flex: 1; font-size: $font-subhead; color: $color-label; }
}

// --- Category Scroll ---
.cat-scroll {
	white-space: nowrap;
	margin: 0 -#{$spacing-xl};
	padding: 0 $spacing-xl;
	margin-bottom: $spacing-md;

	.cat-bar {
		display: inline-flex;
		gap: $spacing-sm;
	}

	.cat-pill {
		display: inline-block;
		padding: $spacing-xs $spacing-lg;
		border-radius: $radius-full;
		background: $color-fill;
		font-size: $font-footnote;
		color: $color-label-tertiary;
		font-weight: 500;
		transition: all 0.25s cubic-bezier(0.25, 0.46, 0.45, 0.94);

		&.active { background: $color-label; color: #fff; }
		&:active { transform: scale(0.94); }
	}
}

// --- Food List ---
.food-list {
	max-height: 400rpx;
	overflow-y: auto;

	.food-row {
		display: flex;
		align-items: center;
		padding: $spacing-md 0;
		position: relative;

		&::after {
			content: '';
			position: absolute;
			bottom: 0;
			left: 0;
			right: 0;
			height: 0.5rpx;
			background: $color-separator;
		}
		&:last-child::after { display: none; }
		&:active { opacity: 0.7; }

		.food-info {
			flex: 1;
			.food-name { display: block; font-size: $font-body; font-weight: 500; color: $color-label; }
			.food-detail { display: block; font-size: $font-caption2; color: $color-label-quaternary; margin-top: 4rpx; }
		}
		.food-add-btn {
			width: 52rpx;
			height: 52rpx;
			border-radius: 50%;
			background: $color-green-light;
			display: flex;
			align-items: center;
			justify-content: center;
			flex-shrink: 0;
			@include press-effect;

			.food-add-icon {
				font-size: 32rpx;
				color: $color-green;
				font-weight: 400;
				line-height: 1;
			}
		}
	}
}

.custom-link {
	text-align: center;
	padding: $spacing-md 0 4rpx;
	font-size: $font-subhead;
	color: $color-primary;
	font-weight: 500;
	&:active { opacity: 0.5; }
}

// --- Selected Foods ---
.selected-row {
	display: flex;
	align-items: center;
	padding: $spacing-md 0;
	position: relative;

	&::after {
		content: '';
		position: absolute;
		bottom: 0;
		left: 0;
		right: 0;
		height: 0.5rpx;
		background: $color-separator;
	}

	.sf-name { flex: 1; font-size: $font-body; color: $color-label; font-weight: 500; }
	.sf-cal { font-size: $font-footnote; color: $color-label-quaternary; margin-right: $spacing-md; font-variant-numeric: tabular-nums; }
	.sf-remove {
		width: 40rpx;
		height: 40rpx;
		border-radius: 50%;
		background: $color-red-light;
		display: flex;
		align-items: center;
		justify-content: center;
		@include press-effect;

		.sf-remove-icon {
			font-size: $font-caption2;
			color: $color-red;
		}
	}
}

.total-bar {
	margin-top: $spacing-lg;
	padding-top: $spacing-lg;
	border-top: 0.5rpx solid $color-separator;

	.total-val {
		display: block;
		font-size: $font-title2;
		font-weight: 700;
		color: $color-label;
		letter-spacing: -1rpx;
		font-variant-numeric: tabular-nums;
	}
	.total-detail {
		display: block;
		font-size: $font-caption1;
		color: $color-label-quaternary;
		margin-top: 4rpx;
	}
}

.notes-area {
	width: 100%;
	height: 120rpx;
	font-size: $font-subhead;
	color: $color-label;
	line-height: 1.6;
}

// --- Bottom Action ---
.bottom-action {
	@include bottom-action-bar;
	.save-btn { @include primary-button; }
}

// --- Custom Food Sheet ---
.mask {
	position: fixed;
	top: 0; left: 0; right: 0; bottom: 0;
	background: rgba(0, 0, 0, 0.35);
	z-index: 100;
}

.custom-sheet {
	position: fixed;
	bottom: 0;
	left: 0;
	right: 0;
	background: $color-bg-elevated;
	border-radius: $radius-xl $radius-xl 0 0;
	padding: $spacing-md $spacing-xl $spacing-2xl;
	padding-bottom: calc(#{$spacing-2xl} + env(safe-area-inset-bottom));
	z-index: 101;

	.sheet-handle {
		width: 72rpx;
		height: 8rpx;
		background: $color-fill-tertiary;
		border-radius: 4rpx;
		margin: 0 auto $spacing-lg;
	}

	.sheet-title {
		display: block;
		font-size: $font-title3;
		font-weight: 700;
		color: $color-label;
		text-align: center;
		margin-bottom: $spacing-xl;
	}

	.sheet-field {
		margin-bottom: $spacing-md;

		.sheet-label {
			display: block;
			font-size: $font-footnote;
			color: $color-label-quaternary;
			font-weight: 500;
			margin-bottom: $spacing-xs;
		}
		.sheet-input-wrap { @include input-field; }
		.sheet-input { flex: 1; font-size: $font-body; color: $color-label; }
	}

	.sheet-save {
		@include primary-button;
		margin-top: $spacing-lg;
	}
}
</style>
