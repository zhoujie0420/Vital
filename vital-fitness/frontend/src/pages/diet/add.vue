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
				<text class="search-icon">S</text>
				<input class="search-input" v-model="keyword" placeholder="搜索食物名称" @input="searchFoods" />
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
				<view v-for="f in foodList" :key="f.id" class="food-row" @tap="openPortion(f)">
					<view class="food-info">
						<text class="food-name">{{ f.name }}</text>
						<text class="food-detail">{{ f.calories }} kcal/100g</text>
					</view>
					<view class="food-add-btn">
						<text class="food-add-icon">+</text>
					</view>
				</view>
			</view>

			<view class="food-actions">
				<view class="custom-link" @tap="showCustom = true">
					<text>+ 自定义食物</text>
				</view>
				<view class="ai-recognition-btn" @tap="startRecognition">
					<text class="ai-btn-icon">📷</text>
					<text class="ai-btn-text">拍照识别</text>
				</view>
			</view>
		</view>

		<!-- AI 识别加载遮罩 -->
		<view class="recognition-overlay" v-if="recognizing">
			<view class="recognition-loading">
				<view class="recognition-spinner"></view>
				<text class="recognition-loading-text">AI 识别中...</text>
			</view>
		</view>

		<!-- 已选食物 -->
		<view class="card" v-if="selectedFoods.length > 0">
			<text class="card-label">已选 · {{ selectedFoods.length }} 项</text>
			<view v-for="(sf, i) in selectedFoods" :key="i" class="selected-row">
				<view class="sf-info">
					<text class="sf-name">{{ sf.name }}</text>
					<text class="sf-portion">{{ sf.grams }}g</text>
				</view>
				<view class="sf-right">
					<text class="sf-cal">{{ sf.calories }} kcal</text>
					<view class="sf-remove" @tap="selectedFoods.splice(i, 1)">
						<text class="sf-remove-icon">✕</text>
					</view>
				</view>
			</view>

			<!-- 汇总 -->
			<view class="total-bar">
				<view class="total-main">
					<text class="total-val">{{ totalCalories }}</text>
					<text class="total-unit">千卡</text>
				</view>
				<view class="total-macros">
					<text class="total-macro">蛋白 {{ totalProtein }}g</text>
					<text class="total-sep">·</text>
					<text class="total-macro">碳水 {{ totalCarbs }}g</text>
					<text class="total-sep">·</text>
					<text class="total-macro">脂肪 {{ totalFat }}g</text>
				</view>
			</view>
		</view>

		<view class="bottom-action" v-if="selectedFoods.length > 0">
			<view class="save-btn" :class="{ loading: saving }" @tap="save">
				<text>{{ saving ? '保存中...' : '保存记录' }}</text>
			</view>
		</view>

		<!-- AI 识别结果弹窗 -->
		<view class="mask" v-if="showRecognition" @tap="showRecognition = false"></view>
		<view class="recognition-sheet" v-if="showRecognition">
			<view class="sheet-handle"></view>
			<text class="sheet-title">识别结果</text>
			<text class="sheet-hint">可调整份量，勾选后添加</text>

			<scroll-view scroll-y class="recognition-list">
				<view v-for="(food, idx) in recognizedFoods" :key="idx" class="recognition-item">
					<view class="recognition-item-header">
						<view class="recognition-checkbox" :class="{ checked: food.selected }" @tap="food.selected = !food.selected">
							<text v-if="food.selected" class="checkbox-icon">✓</text>
						</view>
						<view class="recognition-item-info">
							<text class="recognition-food-name">{{ food.name }}</text>
							<view class="confidence-indicator" v-if="food.confidence < 0.6">
								<text class="confidence-warning">结果仅供参考</text>
							</view>
						</view>
						<view class="recognition-delete" @tap="recognizedFoods.splice(idx, 1)">
							<text class="recognition-delete-icon">✕</text>
						</view>
					</view>

					<view class="recognition-item-body">
						<view class="recognition-grams-row">
							<text class="recognition-grams-label">份量</text>
							<view class="recognition-grams-input-wrap">
								<input type="digit" v-model="food.estimated_grams" class="recognition-grams-input" placeholder="100" />
								<text class="recognition-grams-suffix">克</text>
							</view>
						</view>
						<view class="recognition-nutrition-preview">
							<text class="recognition-nutrition-item">{{ calcRecognitionCal(food) }} kcal</text>
							<text class="recognition-nutrition-sep">·</text>
							<text class="recognition-nutrition-item">蛋白 {{ calcRecognitionProtein(food) }}g</text>
							<text class="recognition-nutrition-sep">·</text>
							<text class="recognition-nutrition-item">碳水 {{ calcRecognitionCarbs(food) }}g</text>
							<text class="recognition-nutrition-sep">·</text>
							<text class="recognition-nutrition-item">脂肪 {{ calcRecognitionFat(food) }}g</text>
						</view>
					</view>
				</view>
			</scroll-view>

			<view class="recognition-confirm" @tap="confirmRecognition">
				<text>全部添加</text>
			</view>
		</view>

		<!-- 份量选择弹窗 -->
		<view class="mask" v-if="showPortion" @tap="showPortion = false"></view>
		<view class="portion-sheet" v-if="showPortion">
			<view class="sheet-handle"></view>
			<text class="portion-food-name">{{ portionFood.name }}</text>
			<text class="portion-per100">每100g：{{ portionFood.calories }} kcal · 蛋白{{ portionFood.protein }}g · 碳水{{ portionFood.carbs }}g · 脂肪{{ portionFood.fat }}g</text>

			<!-- 快捷份量 -->
			<view class="portion-quick">
				<text v-for="g in quickGrams" :key="g" class="portion-chip"
					:class="{ active: portionGrams === String(g) }" @tap="portionGrams = String(g)">
					{{ g }}g
				</text>
			</view>

			<!-- 自定义克数 -->
			<view class="portion-input-row">
				<text class="portion-input-label">份量</text>
				<view class="portion-input-wrap">
					<input type="digit" v-model="portionGrams" class="portion-input" placeholder="100" />
					<text class="portion-input-suffix">克</text>
				</view>
			</view>

			<!-- 实时计算预览 -->
			<view class="portion-preview" v-if="portionGrams">
				<view class="preview-cal">
					<text class="preview-cal-val">{{ calcCal }}</text>
					<text class="preview-cal-unit">kcal</text>
				</view>
				<view class="preview-macros">
					<text class="preview-macro">蛋白 {{ calcProtein }}g</text>
					<text class="preview-macro">碳水 {{ calcCarbs }}g</text>
					<text class="preview-macro">脂肪 {{ calcFat }}g</text>
				</view>
			</view>

			<view class="portion-confirm" @tap="confirmPortion">
				<text>添加</text>
			</view>
		</view>

		<!-- 自定义食物弹窗 -->
		<view class="mask" v-if="showCustom" @tap="showCustom = false"></view>
		<view class="custom-sheet" v-if="showCustom">
			<view class="sheet-handle"></view>
			<text class="sheet-title">自定义食物</text>
			<text class="sheet-hint">营养数据按每 100g 填写</text>
			<view class="sheet-field">
				<text class="sheet-label">名称</text>
				<view class="sheet-input-wrap">
					<input v-model="customFood.name" placeholder="食物名称" class="sheet-input" />
				</view>
			</view>
			<view class="sheet-row">
				<view class="sheet-field sheet-half">
					<text class="sheet-label">热量 (kcal)</text>
					<view class="sheet-input-wrap">
						<input v-model="customFood.calories" type="digit" placeholder="0" class="sheet-input" />
					</view>
				</view>
				<view class="sheet-field sheet-half">
					<text class="sheet-label">蛋白质 (g)</text>
					<view class="sheet-input-wrap">
						<input v-model="customFood.protein" type="digit" placeholder="0" class="sheet-input" />
					</view>
				</view>
			</view>
			<view class="sheet-row">
				<view class="sheet-field sheet-half">
					<text class="sheet-label">碳水 (g)</text>
					<view class="sheet-input-wrap">
						<input v-model="customFood.carbs" type="digit" placeholder="0" class="sheet-input" />
					</view>
				</view>
				<view class="sheet-field sheet-half">
					<text class="sheet-label">脂肪 (g)</text>
					<view class="sheet-input-wrap">
						<input v-model="customFood.fat" type="digit" placeholder="0" class="sheet-input" />
					</view>
				</view>
			</view>
			<view class="sheet-save" @tap="saveCustomFood">
				<text>保存</text>
			</view>
		</view>
	</view>
</template>

<script>
	import { getFoods, createFood, createDietRecord, recognizeFood } from '../../api/diet'

	export default {
		data() {
			return {
				saving: false, showCustom: false, showPortion: false,
				showRecognition: false, recognizing: false, recognizedFoods: [], recognitionError: '',
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
				// 份量弹窗
				portionFood: {},
				portionGrams: '100',
				quickGrams: [50, 100, 150, 200, 300],
				customFood: { name: '', calories: '', protein: '', carbs: '', fat: '' }
			}
		},
		computed: {
			topPadding() {
				const app = getApp()
				return (app.globalData?.customBarHeight || 88) + 8
			},
			// 份量计算（按比例）
			calcCal() {
				const g = parseFloat(this.portionGrams) || 0
				return ((this.portionFood.calories || 0) * g / 100).toFixed(0)
			},
			calcProtein() {
				const g = parseFloat(this.portionGrams) || 0
				return ((this.portionFood.protein || 0) * g / 100).toFixed(1)
			},
			calcCarbs() {
				const g = parseFloat(this.portionGrams) || 0
				return ((this.portionFood.carbs || 0) * g / 100).toFixed(1)
			},
			calcFat() {
				const g = parseFloat(this.portionGrams) || 0
				return ((this.portionFood.fat || 0) * g / 100).toFixed(1)
			},
			// 已选汇总
			totalCalories() { return this.selectedFoods.reduce((s, f) => s + f.calories, 0).toFixed(0) },
			totalProtein() { return this.selectedFoods.reduce((s, f) => s + f.protein, 0).toFixed(1) },
			totalCarbs() { return this.selectedFoods.reduce((s, f) => s + f.carbs, 0).toFixed(1) },
			totalFat() { return this.selectedFoods.reduce((s, f) => s + f.fat, 0).toFixed(1) }
		},
		onLoad(opts) {
			if (opts.meal) this.mealType = opts.meal
			this.searchFoods()
		},
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
			openPortion(f) {
				this.portionFood = { ...f }
				this.portionGrams = '100'
				this.showPortion = true
			},
			confirmPortion() {
				const g = parseFloat(this.portionGrams) || 0
				if (g <= 0) {
					uni.showToast({ title: '请输入份量', icon: 'none' }); return
				}
				const ratio = g / 100
				this.selectedFoods.push({
					name: this.portionFood.name,
					grams: g,
					calories: parseFloat((this.portionFood.calories * ratio).toFixed(1)),
					protein: parseFloat((this.portionFood.protein * ratio).toFixed(1)),
					carbs: parseFloat((this.portionFood.carbs * ratio).toFixed(1)),
					fat: parseFloat((this.portionFood.fat * ratio).toFixed(1))
				})
				this.showPortion = false
			},
			async saveCustomFood() {
				if (!this.customFood.name || !this.customFood.calories) {
					uni.showToast({ title: '请填写名称和热量', icon: 'none' }); return
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
						food_items: JSON.stringify(this.selectedFoods.map(f => ({ name: f.name, grams: f.grams, calories: f.calories }))),
						total_calories: parseFloat(this.totalCalories),
						protein: parseFloat(this.totalProtein),
						carbs: parseFloat(this.totalCarbs),
						fat: parseFloat(this.totalFat),
						notes: ''
					})
					uni.showToast({ title: '保存成功', icon: 'success' })
					setTimeout(() => uni.navigateBack(), 1000)
				} catch (e) {} finally { this.saving = false }
			},
			startRecognition() {
				uni.chooseImage({
					count: 1,
					sizeType: ['compressed'],
					sourceType: ['album', 'camera'],
					success: (res) => {
						const tempFilePath = res.tempFilePaths[0]
						this.readFileAsBase64(tempFilePath)
					},
					fail: (err) => {
						if (err.errMsg && (err.errMsg.indexOf('auth deny') !== -1 || err.errMsg.indexOf('authorize') !== -1)) {
							uni.showModal({
								title: '需要相机权限',
								content: '请前往设置 → 隐私 → 相机，开启相机权限',
								confirmText: '去设置',
								success: (modalRes) => {
									if (modalRes.confirm) {
										uni.openSetting()
									}
								}
							})
						}
					}
				})
			},
			readFileAsBase64(filePath) {
				// #ifdef MP-WEIXIN
				const fs = uni.getFileSystemManager()
				fs.readFile({
					filePath: filePath,
					encoding: 'base64',
					success: (res) => {
						this.processBase64Image(res.data)
					},
					fail: () => {
						uni.showToast({ title: '读取图片失败', icon: 'none' })
					}
				})
				// #endif
				// #ifdef H5
				uni.request({
					url: filePath,
					responseType: 'arraybuffer',
					success: (res) => {
						const base64 = uni.arrayBufferToBase64(res.data)
						this.processBase64Image(base64)
					},
					fail: () => {
						uni.showToast({ title: '读取图片失败', icon: 'none' })
					}
				})
				// #endif
			},
			async processBase64Image(base64Data) {
				// Check size: base64 length * 3/4 gives approximate byte size
				const sizeInBytes = base64Data.length * 3 / 4
				const sizeInMB = sizeInBytes / (1024 * 1024)
				if (sizeInMB > 2) {
					uni.showToast({ title: '图片过大，请重新选择', icon: 'none' })
					return
				}

				this.recognizing = true
				this.recognitionError = ''
				try {
					const res = await recognizeFood(base64Data)
					const foods = (res.data && res.data.foods) || []
					if (foods.length === 0) {
						uni.showToast({ title: '未能识别出食物，请尝试重新拍照', icon: 'none' })
						return
					}
					this.recognizedFoods = foods.map(f => ({ ...f, selected: true }))
					this.showRecognition = true
				} catch (e) {
					const msg = (e && e.message) || 'AI 识别失败，请重试'
					this.recognitionError = msg
					uni.showToast({ title: msg, icon: 'none' })
				} finally {
					this.recognizing = false
				}
			},
			calcRecognitionCal(food) {
				const g = parseFloat(food.estimated_grams) || 0
				return ((food.calories_per_100g || 0) * g / 100).toFixed(0)
			},
			calcRecognitionProtein(food) {
				const g = parseFloat(food.estimated_grams) || 0
				return ((food.protein_per_100g || 0) * g / 100).toFixed(1)
			},
			calcRecognitionCarbs(food) {
				const g = parseFloat(food.estimated_grams) || 0
				return ((food.carbs_per_100g || 0) * g / 100).toFixed(1)
			},
			calcRecognitionFat(food) {
				const g = parseFloat(food.estimated_grams) || 0
				return ((food.fat_per_100g || 0) * g / 100).toFixed(1)
			},
			convertToSelectedFood(recognizedFood) {
				return {
					name: recognizedFood.name,
					grams: parseFloat(recognizedFood.estimated_grams) || 0,
					calories: parseFloat((recognizedFood.calories_per_100g * (parseFloat(recognizedFood.estimated_grams) || 0) / 100).toFixed(1)),
					protein: parseFloat((recognizedFood.protein_per_100g * (parseFloat(recognizedFood.estimated_grams) || 0) / 100).toFixed(1)),
					carbs: parseFloat((recognizedFood.carbs_per_100g * (parseFloat(recognizedFood.estimated_grams) || 0) / 100).toFixed(1)),
					fat: parseFloat((recognizedFood.fat_per_100g * (parseFloat(recognizedFood.estimated_grams) || 0) / 100).toFixed(1))
				}
			},
			confirmRecognition() {
				const selected = this.recognizedFoods.filter(f => f.selected)
				const converted = selected.map(f => this.convertToSelectedFood(f))
				this.selectedFoods.push(...converted)
				this.showRecognition = false
				this.recognizedFoods = []
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

	.search-icon { font-size: 24rpx; margin-right: $spacing-sm; color: $color-label-quaternary; font-weight: 600; }
	.search-input { flex: 1; font-size: $font-subhead; color: $color-label; }
}

// --- Category Scroll ---
.cat-scroll {
	white-space: nowrap;
	margin: 0 -#{$spacing-xl};
	padding: 0 $spacing-xl;
	margin-bottom: $spacing-md;

	.cat-bar { display: inline-flex; gap: $spacing-sm; }

	.cat-pill {
		display: inline-block;
		padding: $spacing-xs $spacing-lg;
		border-radius: $radius-full;
		background: $color-fill;
		font-size: $font-footnote;
		color: $color-label-tertiary;
		font-weight: 500;
		transition: all 0.25s cubic-bezier(0.16, 1, 0.3, 1);

		&.active { background: $color-label; color: #fff; }
		&:active { transform: scale(0.94); }
	}
}

// --- Food List ---
.food-list {
	max-height: 480rpx;
	overflow-y: auto;

	.food-row {
		display: flex;
		align-items: center;
		padding: $spacing-md 0;
		position: relative;
		@include press-effect;

		&::after {
			content: '';
			position: absolute;
			bottom: 0; left: 0; right: 0;
			height: 0.5rpx;
			background: $color-separator;
		}
		&:last-child::after { display: none; }

		.food-info {
			flex: 1;
			.food-name { display: block; font-size: $font-body; font-weight: 500; color: $color-label; }
			.food-detail { display: block; font-size: $font-caption2; color: $color-label-quaternary; margin-top: 4rpx; font-variant-numeric: tabular-nums; }
		}
		.food-add-btn {
			width: 52rpx; height: 52rpx;
			border-radius: 50%;
			background: $color-green-light;
			display: flex; align-items: center; justify-content: center;
			flex-shrink: 0;

			.food-add-icon { font-size: 32rpx; color: $color-green; font-weight: 400; line-height: 1; }
		}
	}
}

// --- Selected Foods ---
.selected-row {
	display: flex;
	align-items: center;
	justify-content: space-between;
	padding: $spacing-md 0;
	position: relative;

	&::after {
		content: '';
		position: absolute;
		bottom: 0; left: 0; right: 0;
		height: 0.5rpx;
		background: $color-separator;
	}
	&:last-child::after { display: none; }

	.sf-info {
		flex: 1;
		.sf-name { display: block; font-size: $font-body; color: $color-label; font-weight: 500; }
		.sf-portion { display: block; font-size: $font-caption2; color: $color-label-quaternary; margin-top: 2rpx; }
	}
	.sf-right {
		display: flex;
		align-items: center;
		gap: $spacing-md;

		.sf-cal { font-size: $font-subhead; color: $color-accent; font-weight: 600; font-variant-numeric: tabular-nums; }
	}
	.sf-remove {
		width: 40rpx; height: 40rpx;
		border-radius: 50%;
		background: $color-red-light;
		display: flex; align-items: center; justify-content: center;
		@include press-effect;

		.sf-remove-icon { font-size: $font-caption2; color: $color-red; }
	}
}

// --- Total Bar ---
.total-bar {
	margin-top: $spacing-lg;
	padding-top: $spacing-lg;
	border-top: 0.5rpx solid $color-separator;

	.total-main {
		display: flex;
		align-items: baseline;
		gap: $spacing-xs;
	}
	.total-val {
		font-size: $font-title2;
		font-weight: 700;
		color: $color-label;
		letter-spacing: -1rpx;
		font-variant-numeric: tabular-nums;
	}
	.total-unit {
		font-size: $font-footnote;
		color: $color-label-quaternary;
	}
	.total-macros {
		display: flex;
		align-items: center;
		gap: $spacing-xs;
		margin-top: 4rpx;
	}
	.total-macro { font-size: $font-caption1; color: $color-label-quaternary; font-weight: 500; }
	.total-sep { font-size: $font-caption1; color: $color-separator-opaque; }
}

// --- Bottom Action ---
.bottom-action {
	@include bottom-action-bar;
	.save-btn { @include primary-button; }
}

// --- Mask ---
.mask {
	position: fixed;
	top: 0; left: 0; right: 0; bottom: 0;
	background: rgba(0, 0, 0, 0.35);
	z-index: 100;
}

// --- Portion Sheet ---
.portion-sheet {
	position: fixed;
	bottom: 0; left: 0; right: 0;
	background: $color-bg-elevated;
	border-radius: $radius-xl $radius-xl 0 0;
	padding: $spacing-md $spacing-xl $spacing-2xl;
	padding-bottom: calc(#{$spacing-2xl} + env(safe-area-inset-bottom));
	z-index: 101;

	.sheet-handle {
		width: 72rpx; height: 8rpx;
		background: $color-fill-tertiary;
		border-radius: 4rpx;
		margin: 0 auto $spacing-lg;
	}
}

.portion-food-name {
	display: block;
	font-size: $font-title3;
	font-weight: 700;
	color: $color-label;
	text-align: center;
	margin-bottom: $spacing-xs;
}

.portion-per100 {
	display: block;
	font-size: $font-caption1;
	color: $color-label-quaternary;
	text-align: center;
	margin-bottom: $spacing-xl;
	line-height: 1.5;
}

// --- Quick Grams ---
.portion-quick {
	display: flex;
	gap: $spacing-sm;
	margin-bottom: $spacing-lg;

	.portion-chip {
		flex: 1;
		text-align: center;
		padding: $spacing-md 0;
		background: $color-fill;
		border-radius: $radius-md;
		font-size: $font-subhead;
		font-weight: 600;
		color: $color-label-tertiary;
		font-variant-numeric: tabular-nums;
		transition: all 0.2s cubic-bezier(0.16, 1, 0.3, 1);

		&.active {
			background: $color-primary;
			color: #fff;
			box-shadow: 0 4rpx 12rpx rgba(16, 185, 129, 0.2);
		}
		&:active { transform: scale(0.94); }
	}
}

// --- Portion Input ---
.portion-input-row {
	display: flex;
	align-items: center;
	gap: $spacing-md;
	margin-bottom: $spacing-lg;

	.portion-input-label {
		font-size: $font-body;
		color: $color-label;
		font-weight: 500;
		flex-shrink: 0;
	}
}

.portion-input-wrap {
	flex: 1;
	display: flex;
	align-items: center;
	background: $color-fill;
	border-radius: $radius-md;
	padding: $spacing-md $spacing-lg;
	border: 1.5rpx solid transparent;
	transition: all 0.2s ease;

	&:focus-within {
		background: rgba(16, 185, 129, 0.04);
		border-color: rgba(16, 185, 129, 0.3);
	}

	.portion-input {
		flex: 1;
		font-size: $font-headline;
		font-weight: 700;
		color: $color-label;
		text-align: center;
		font-variant-numeric: tabular-nums;
	}
	.portion-input-suffix {
		font-size: $font-subhead;
		color: $color-label-quaternary;
		font-weight: 500;
	}
}

// --- Preview ---
.portion-preview {
	background: $color-fill;
	border-radius: $radius-md;
	padding: $spacing-lg;
	margin-bottom: $spacing-lg;
	text-align: center;

	.preview-cal {
		display: flex;
		align-items: baseline;
		justify-content: center;
		gap: $spacing-xs;
		margin-bottom: $spacing-xs;
	}
	.preview-cal-val {
		font-size: 48rpx;
		font-weight: 700;
		color: $color-primary;
		font-variant-numeric: tabular-nums;
	}
	.preview-cal-unit {
		font-size: $font-footnote;
		color: $color-label-quaternary;
	}
	.preview-macros {
		display: flex;
		justify-content: center;
		gap: $spacing-lg;
	}
	.preview-macro {
		font-size: $font-caption1;
		color: $color-label-tertiary;
		font-weight: 500;
	}
}

.portion-confirm {
	@include primary-button;
}

// --- Custom Food Sheet ---
.custom-sheet {
	position: fixed;
	bottom: 0; left: 0; right: 0;
	background: $color-bg-elevated;
	border-radius: $radius-xl $radius-xl 0 0;
	padding: $spacing-md $spacing-xl $spacing-2xl;
	padding-bottom: calc(#{$spacing-2xl} + env(safe-area-inset-bottom));
	z-index: 101;

	.sheet-handle {
		width: 72rpx; height: 8rpx;
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
		margin-bottom: $spacing-xs;
	}
	.sheet-hint {
		display: block;
		font-size: $font-caption1;
		color: $color-label-quaternary;
		text-align: center;
		margin-bottom: $spacing-xl;
	}

	.sheet-row {
		display: flex;
		gap: $spacing-sm;
	}
	.sheet-half {
		flex: 1;
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

// --- Food Actions (custom link + AI button) ---
.food-actions {
	display: flex;
	align-items: center;
	justify-content: center;
	gap: $spacing-xl;
	padding: $spacing-md 0 4rpx;
}

.custom-link {
	font-size: $font-subhead;
	color: $color-primary;
	font-weight: 500;
	&:active { opacity: 0.5; }
}

// --- AI Recognition Button ---
.ai-recognition-btn {
	display: flex;
	align-items: center;
	gap: $spacing-xs;
	padding: $spacing-sm $spacing-lg;
	background: $color-primary-light;
	border-radius: $radius-full;
	@include press-effect;

	.ai-btn-icon {
		font-size: 28rpx;
	}
	.ai-btn-text {
		font-size: $font-subhead;
		color: $color-primary;
		font-weight: 600;
	}
}

// --- Recognition Loading Overlay ---
.recognition-overlay {
	position: fixed;
	top: 0; left: 0; right: 0; bottom: 0;
	background: rgba(0, 0, 0, 0.5);
	z-index: 200;
	display: flex;
	align-items: center;
	justify-content: center;

	.recognition-loading {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: $spacing-lg;
		background: $color-bg-elevated;
		border-radius: $radius-xl;
		padding: $spacing-3xl $spacing-3xl;
		box-shadow: $shadow-float;
	}

	.recognition-spinner {
		width: 64rpx;
		height: 64rpx;
		border: 4rpx solid $color-fill-tertiary;
		border-top-color: $color-primary;
		border-radius: 50%;
		animation: spin 0.8s linear infinite;
	}

	.recognition-loading-text {
		font-size: $font-body;
		color: $color-label;
		font-weight: 600;
	}
}

@keyframes spin {
	to { transform: rotate(360deg); }
}

// --- Recognition Result Sheet ---
.recognition-sheet {
	position: fixed;
	bottom: 0; left: 0; right: 0;
	background: $color-bg-elevated;
	border-radius: $radius-xl $radius-xl 0 0;
	padding: $spacing-md $spacing-xl $spacing-2xl;
	padding-bottom: calc(#{$spacing-2xl} + env(safe-area-inset-bottom));
	z-index: 101;
	max-height: 75vh;
	display: flex;
	flex-direction: column;

	.sheet-handle {
		width: 72rpx; height: 8rpx;
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
		margin-bottom: $spacing-xs;
	}

	.sheet-hint {
		display: block;
		font-size: $font-caption1;
		color: $color-label-quaternary;
		text-align: center;
		margin-bottom: $spacing-lg;
	}
}

.recognition-list {
	flex: 1;
	max-height: 50vh;
	margin-bottom: $spacing-lg;
}

.recognition-item {
	background: $color-fill;
	border-radius: $radius-md;
	padding: $spacing-lg;
	margin-bottom: $spacing-sm;

	.recognition-item-header {
		display: flex;
		align-items: center;
		gap: $spacing-md;
		margin-bottom: $spacing-md;
	}

	.recognition-checkbox {
		width: 44rpx; height: 44rpx;
		border-radius: $radius-sm;
		border: 2rpx solid $color-label-quaternary;
		display: flex;
		align-items: center;
		justify-content: center;
		flex-shrink: 0;
		transition: all 0.2s ease;

		&.checked {
			background: $color-primary;
			border-color: $color-primary;
		}

		.checkbox-icon {
			font-size: 24rpx;
			color: #fff;
			font-weight: 700;
		}
	}

	.recognition-item-info {
		flex: 1;
	}

	.recognition-food-name {
		display: block;
		font-size: $font-body;
		font-weight: 600;
		color: $color-label;
	}

	.confidence-indicator {
		margin-top: 4rpx;
	}

	.confidence-warning {
		font-size: $font-caption2;
		color: $color-orange;
		font-weight: 500;
		background: $color-orange-light;
		padding: 2rpx $spacing-sm;
		border-radius: $radius-sm;
	}

	.recognition-delete {
		width: 40rpx; height: 40rpx;
		border-radius: 50%;
		background: $color-red-light;
		display: flex;
		align-items: center;
		justify-content: center;
		flex-shrink: 0;
		@include press-effect;

		.recognition-delete-icon {
			font-size: $font-caption2;
			color: $color-red;
		}
	}

	.recognition-item-body {
		padding-left: 60rpx;
	}

	.recognition-grams-row {
		display: flex;
		align-items: center;
		gap: $spacing-md;
		margin-bottom: $spacing-sm;

		.recognition-grams-label {
			font-size: $font-footnote;
			color: $color-label-tertiary;
			font-weight: 500;
			flex-shrink: 0;
		}

		.recognition-grams-input-wrap {
			display: flex;
			align-items: center;
			background: $color-bg-elevated;
			border-radius: $radius-sm;
			padding: $spacing-sm $spacing-md;
			border: 1.5rpx solid transparent;
			transition: all 0.2s ease;
			flex: 1;
			max-width: 240rpx;

			&:focus-within {
				border-color: rgba(16, 185, 129, 0.3);
			}

			.recognition-grams-input {
				flex: 1;
				font-size: $font-subhead;
				font-weight: 600;
				color: $color-label;
				text-align: center;
				font-variant-numeric: tabular-nums;
			}

			.recognition-grams-suffix {
				font-size: $font-caption1;
				color: $color-label-quaternary;
				font-weight: 500;
			}
		}
	}

	.recognition-nutrition-preview {
		display: flex;
		flex-wrap: wrap;
		align-items: center;
		gap: $spacing-xs;

		.recognition-nutrition-item {
			font-size: $font-caption1;
			color: $color-label-tertiary;
			font-weight: 500;
			font-variant-numeric: tabular-nums;
		}

		.recognition-nutrition-sep {
			font-size: $font-caption1;
			color: $color-separator-opaque;
		}
	}
}

.recognition-confirm {
	@include primary-button;
	flex-shrink: 0;
}
</style>
