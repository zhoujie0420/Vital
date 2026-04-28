<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="page-header">
			<text class="page-title">饮食记录</text>
			<text class="header-action" @tap="goToPlan">📋 方案</text>
		</view>

		<!-- 今日概览 — 有方案时显示饼图 -->
		<view class="today-card" v-if="hasPlan">
			<view class="pie-area">
				<view class="pie-chart" :style="pieStyle">
					<view class="pie-center">
						<text class="pie-val">{{ todayCal }}</text>
						<text class="pie-unit">kcal</text>
					</view>
				</view>
			</view>
			<view class="legend-area">
				<view class="legend-row">
					<view class="legend-dot lg-teal"></view>
					<text class="legend-name">蛋白质</text>
					<text class="legend-val">{{ todayProtein }}<text class="legend-target"> / {{ targetProtein }}g</text></text>
				</view>
				<view class="legend-row">
					<view class="legend-dot lg-orange"></view>
					<text class="legend-name">碳水</text>
					<text class="legend-val">{{ todayCarbs }}<text class="legend-target"> / {{ targetCarbs }}g</text></text>
				</view>
				<view class="legend-row">
					<view class="legend-dot lg-yellow"></view>
					<text class="legend-name">脂肪</text>
					<text class="legend-val">{{ todayFat }}<text class="legend-target"> / {{ targetFat }}g</text></text>
				</view>
			</view>
		</view>

		<!-- 无方案时显示纯数字 + 引导 -->
		<view class="today-card today-simple" v-else>
			<view class="simple-stats">
				<view class="simple-main">
					<text class="simple-val">{{ todayCal }}</text>
					<text class="simple-unit">千卡</text>
				</view>
				<text class="simple-label">今日已摄入</text>
				<view class="simple-macros">
					<text class="simple-macro">蛋白{{ todayProtein }}g</text>
					<text class="simple-sep">·</text>
					<text class="simple-macro">碳水{{ todayCarbs }}g</text>
					<text class="simple-sep">·</text>
					<text class="simple-macro">脂肪{{ todayFat }}g</text>
				</view>
			</view>
			<view class="plan-hint" @tap="goToPlan">
				<text class="plan-hint-text">设置饮食目标后可查看进度</text>
				<text class="plan-hint-link">去设置 ›</text>
			</view>
		</view>

		<!-- 按餐次分区 -->
		<view class="meal-card" v-for="meal in mealSections" :key="meal.type">
			<view class="meal-header">
				<view class="meal-left">
					<view class="meal-dot" :class="'dot-' + meal.type"></view>
					<text class="meal-name">{{ meal.name }}</text>
				</view>
				<view class="meal-right">
					<text class="meal-cal" v-if="getMealCal(meal.type) > 0">{{ getMealCal(meal.type) }} kcal</text>
					<view class="meal-add" @tap="addDiet(meal.type)"><text>+</text></view>
				</view>
			</view>
			<!-- 该餐进度条 -->
			<view class="meal-progress" v-if="getMealCal(meal.type) > 0">
				<view class="meal-progress-fill" :class="'fill-' + meal.type"
					:style="{ width: macroPct(getMealCal(meal.type), targetCal) }"></view>
			</view>
			<!-- 该餐记录 -->
			<view v-if="getMealRecords(meal.type).length > 0" class="meal-items">
				<view v-for="r in getMealRecords(meal.type)" :key="r.id" class="meal-item" @tap="goToDetail(getRecordDate(r))">
					<text class="mi-cal">{{ r.total_calories }} kcal</text>
					<text class="mi-macro">蛋白{{ r.protein }}g · 碳水{{ r.carbs }}g · 脂肪{{ r.fat }}g</text>
				</view>
			</view>
		</view>

		<!-- 历史记录 -->
		<view class="history-section" v-if="historyDays.length > 0">
			<text class="history-title">历史记录</text>
			<view v-for="(day, i) in historyDays" :key="i" class="history-card" @tap="goToDetail(day.date)">
				<view class="history-top">
					<view class="history-left">
						<text class="history-date">{{ day.label }}</text>
						<text class="history-sub">{{ day.mealCount }} 餐</text>
					</view>
					<view class="history-right">
						<text class="history-cal">{{ day.totalCal }} kcal</text>
						<text class="history-arrow">›</text>
					</view>
				</view>
				<view class="history-macros">
					<text class="hm-item">蛋白{{ day.protein }}g</text>
					<text class="hm-sep">·</text>
					<text class="hm-item">碳水{{ day.carbs }}g</text>
					<text class="hm-sep">·</text>
					<text class="hm-item">脂肪{{ day.fat }}g</text>
				</view>
			</view>
		</view>

		<view v-if="dietList.length === 0 && !loading" class="empty">
			<view class="empty-icon"><text class="empty-icon-letter">D</text></view>
			<text class="empty-title">暂无饮食记录</text>
			<text class="empty-desc">点击每餐旁的 + 开始记录</text>
		</view>
	</view>
</template>

<script>
	import { getDietRecords } from '../../api/diet'
	import { useDietPlanStore } from '../../store/dietPlan'
	import { getTodayTargets } from '../../utils/dietPlanCalc'

	export default {
		data() {
			return {
				dietList: [], page: 1, pageSize: 50, total: 0, loading: false,
				mealSections: [
					{ type: 'breakfast', name: '早餐' },
					{ type: 'lunch', name: '午餐' },
					{ type: 'dinner', name: '晚餐' },
					{ type: 'snack', name: '加餐' }
				]
			}
		},
		computed: {
			topPadding() {
				const app = getApp()
				return (app.globalData?.customBarHeight || 88) + 8
			},
			todayKey() { return new Date().toISOString().split('T')[0] },
			todayRecords() {
				return this.dietList.filter(r => new Date(r.record_date).toISOString().split('T')[0] === this.todayKey)
			},
			todayCal() { return this.todayRecords.reduce((s, r) => s + (r.total_calories || 0), 0).toFixed(0) },
			todayProtein() { return this.todayRecords.reduce((s, r) => s + (r.protein || 0), 0).toFixed(0) },
			todayCarbs() { return this.todayRecords.reduce((s, r) => s + (r.carbs || 0), 0).toFixed(0) },
			todayFat() { return this.todayRecords.reduce((s, r) => s + (r.fat || 0), 0).toFixed(0) },
			// 从饮食规划获取目标
			hasPlan() {
				const store = useDietPlanStore()
				store.init()
				return !!store.activePlan
			},
			planTargets() {
				const store = useDietPlanStore()
				if (store.activePlan) return getTodayTargets(store.activePlan, new Date())
				return null
			},
			targetCal() { return this.planTargets?.calories || 2000 },
			targetProtein() { return this.planTargets?.protein || 120 },
			targetCarbs() { return this.planTargets?.carbs || 250 },
			targetFat() { return this.planTargets?.fat || 65 },
			pieStyle() {
				const p = parseFloat(this.todayProtein) * 4 || 0
				const c = parseFloat(this.todayCarbs) * 4 || 0
				const f = parseFloat(this.todayFat) * 9 || 0
				const total = p + c + f
				if (total === 0) return 'background: #E4E4E7'
				const pPct = (p / total) * 360
				const cPct = (c / total) * 360
				return `background: conic-gradient(#14B8A6 0deg ${pPct}deg, #F97316 ${pPct}deg ${pPct + cPct}deg, #EAB308 ${pPct + cPct}deg 360deg)`
			},
			historyDays() {
				const map = {}
				this.dietList.forEach(r => {
					const key = new Date(r.record_date).toISOString().split('T')[0]
					if (key === this.todayKey) return
					if (!map[key]) map[key] = { date: key, records: [] }
					map[key].records.push(r)
				})
				return Object.values(map).sort((a, b) => b.date.localeCompare(a.date)).slice(0, 10).map(day => {
					const d = new Date(day.date + 'T00:00:00')
					const yesterday = new Date(); yesterday.setDate(yesterday.getDate() - 1)
					const label = d.toDateString() === yesterday.toDateString() ? '昨天' : (d.getMonth()+1) + '月' + d.getDate() + '日'
					return {
						...day, label,
						totalCal: day.records.reduce((s, r) => s + (r.total_calories || 0), 0).toFixed(0),
						protein: day.records.reduce((s, r) => s + (r.protein || 0), 0).toFixed(0),
						carbs: day.records.reduce((s, r) => s + (r.carbs || 0), 0).toFixed(0),
						fat: day.records.reduce((s, r) => s + (r.fat || 0), 0).toFixed(0),
						mealCount: new Set(day.records.map(r => r.meal_type)).size
					}
				})
			}
		},
		onShow() { this.refresh() },
		onPullDownRefresh() { this.refresh().then(() => uni.stopPullDownRefresh()) },
		onReachBottom() { this.loadMore() },
		methods: {
			getMealRecords(type) { return this.todayRecords.filter(r => r.meal_type === type) },
			getMealCal(type) { return this.getMealRecords(type).reduce((s, r) => s + (r.total_calories || 0), 0).toFixed(0) },
			getRecordDate(r) { return new Date(r.record_date).toISOString().split('T')[0] },
			macroPct(val, target) {
				if (!target || target <= 0) return '0%'
				return Math.min(Math.round((parseFloat(val) / target) * 100), 100) + '%'
			},
			async refresh() { this.page = 1; this.dietList = []; await this.load() },
			async loadMore() {
				if (this.loading || this.dietList.length >= this.total) return
				this.page++; await this.load()
			},
			async load() {
				if (this.loading) return
				this.loading = true
				try {
					const res = await getDietRecords({ page: this.page, page_size: this.pageSize })
					const list = res.data || []
					if (this.page === 1) this.dietList = list
					else this.dietList = [...this.dietList, ...list]
					this.total = res.total || 0
				} catch (e) {}
				this.loading = false
			},
			addDiet(mealType) { uni.navigateTo({ url: '/pages/diet/add?meal=' + (mealType || '') }) },
			goToPlan() { uni.navigateTo({ url: '/pages/diet/plan' }) },
			goToDetail(date) { uni.navigateTo({ url: '/pages/diet/detail?date=' + date }) }
		}
	}
</script>

<style lang="scss" scoped>
@import '../../styles/variables.scss';

.page { padding: 0 $spacing-xl; padding-bottom: 40rpx; min-height: 100vh; background: $color-bg; }

.page-header {
	display: flex; justify-content: space-between; align-items: baseline; margin-bottom: $spacing-lg;
	.page-title { @include page-title; }
	.header-action { font-size: $font-subhead; color: $color-primary; font-weight: 600; }
}

// ========== Today Card ==========
.today-card {
	@include card;
	padding: $spacing-xl;
	margin-bottom: $spacing-lg;
	display: flex;
	align-items: center;
	gap: $spacing-xl;
}

// Pie chart
.pie-area {
	flex-shrink: 0;
	text-align: center;

	.pie-sub {
		display: block;
		font-size: $font-caption2;
		color: $color-primary;
		font-weight: 500;
		margin-top: $spacing-sm;
		font-variant-numeric: tabular-nums;
	}
}

.pie-chart {
	width: 160rpx;
	height: 160rpx;
	border-radius: 50%;
	position: relative;
	display: flex;
	align-items: center;
	justify-content: center;

	.pie-center {
		width: 100rpx;
		height: 100rpx;
		border-radius: 50%;
		background: $color-bg-elevated;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		position: relative;
		z-index: 1;

		.pie-val {
			font-size: 32rpx;
			font-weight: 700;
			color: $color-label;
			font-variant-numeric: tabular-nums;
			line-height: 1.1;
		}
		.pie-unit {
			font-size: $font-caption2;
			color: $color-label-quaternary;
		}
	}
}

// Legend
.legend-area {
	flex: 1;
	display: flex;
	flex-direction: column;
	gap: $spacing-md;
}

.legend-row {
	display: flex;
	align-items: center;
	gap: $spacing-sm;

	.legend-dot {
		width: 16rpx;
		height: 16rpx;
		border-radius: 4rpx;
		flex-shrink: 0;
	}
	.lg-teal { background: $color-teal; }
	.lg-orange { background: $color-orange; }
	.lg-yellow { background: $color-yellow; }

	.legend-name {
		font-size: $font-subhead;
		color: $color-label;
		font-weight: 500;
		width: 90rpx;
		flex-shrink: 0;
	}
	.legend-val {
		font-size: $font-subhead;
		color: $color-label;
		font-weight: 700;
		font-variant-numeric: tabular-nums;

		.legend-target {
			font-weight: 400;
			color: $color-label-quaternary;
			font-size: $font-caption1;
		}
	}
}

// Simple mode (no plan)
.today-simple {
	flex-direction: column;
	gap: $spacing-lg;
	text-align: center;

	.simple-stats {
		.simple-main {
			display: flex;
			align-items: baseline;
			justify-content: center;
			gap: $spacing-xs;
		}
		.simple-val {
			font-size: 64rpx;
			font-weight: 700;
			color: $color-label;
			font-variant-numeric: tabular-nums;
			letter-spacing: -2rpx;
		}
		.simple-unit {
			font-size: $font-subhead;
			color: $color-label-quaternary;
		}
		.simple-label {
			display: block;
			font-size: $font-caption1;
			color: $color-label-quaternary;
			margin-top: $spacing-xs;
		}
		.simple-macros {
			display: flex;
			justify-content: center;
			align-items: center;
			gap: $spacing-xs;
			margin-top: $spacing-sm;
		}
		.simple-macro { font-size: $font-caption1; color: $color-label-tertiary; font-weight: 500; }
		.simple-sep { font-size: $font-caption1; color: $color-separator-opaque; }
	}

	.plan-hint {
		display: flex;
		align-items: center;
		justify-content: center;
		gap: $spacing-sm;
		padding: $spacing-md;
		background: $color-fill;
		border-radius: $radius-md;
		@include press-effect;

		.plan-hint-text { font-size: $font-caption1; color: $color-label-quaternary; }
		.plan-hint-link { font-size: $font-caption1; color: $color-primary; font-weight: 600; }
	}
}

// ========== Meal Cards ==========
.meal-card {
	@include card;
	padding: $spacing-lg $spacing-xl;
	margin-bottom: $spacing-sm;
}

.meal-header {
	display: flex;
	justify-content: space-between;
	align-items: center;
}

.meal-left {
	display: flex;
	align-items: center;
	gap: $spacing-sm;
}

.meal-dot {
	width: 16rpx;
	height: 16rpx;
	border-radius: 50%;
}
.dot-breakfast { background: $color-orange; }
.dot-lunch { background: $color-green; }
.dot-dinner { background: $color-teal; }
.dot-snack { background: $color-pink; }

.meal-name {
	font-size: $font-headline;
	font-weight: 600;
	color: $color-label;
}

.meal-right {
	display: flex;
	align-items: center;
	gap: $spacing-sm;
}

.meal-cal {
	font-size: $font-footnote;
	color: $color-label-quaternary;
	font-weight: 600;
	font-variant-numeric: tabular-nums;
}

.meal-add {
	width: 52rpx;
	height: 52rpx;
	border-radius: 50%;
	background: $color-fill;
	display: flex;
	align-items: center;
	justify-content: center;
	@include press-effect;

	text { font-size: 32rpx; color: $color-primary; font-weight: 400; line-height: 1; }
}

// Meal progress bar
.meal-progress {
	height: 8rpx;
	background: $color-fill;
	border-radius: 4rpx;
	margin-top: $spacing-sm;
	overflow: hidden;

	.meal-progress-fill {
		height: 100%;
		border-radius: 4rpx;
		min-width: 4rpx;
		transition: width 0.5s cubic-bezier(0.16, 1, 0.3, 1);
	}
	.fill-breakfast { background: $color-orange; }
	.fill-lunch { background: $color-green; }
	.fill-dinner { background: $color-teal; }
	.fill-snack { background: $color-pink; }
}

.meal-items {
	margin-top: $spacing-md;
	padding-top: $spacing-md;
	border-top: 0.5rpx solid $color-separator;
}

.meal-item {
	padding: $spacing-xs 0;
	@include press-effect;

	.mi-cal {
		display: block;
		font-size: $font-subhead;
		font-weight: 600;
		color: $color-label;
		font-variant-numeric: tabular-nums;
	}
	.mi-macro {
		display: block;
		font-size: $font-caption2;
		color: $color-label-quaternary;
		margin-top: 2rpx;
	}
}

// ========== History ==========
.history-section { margin-top: $spacing-lg; }

.history-title {
	display: block;
	font-size: $font-footnote;
	font-weight: 600;
	color: $color-label-quaternary;
	text-transform: uppercase;
	letter-spacing: 2rpx;
	margin-bottom: $spacing-md;
}

.history-card {
	@include card;
	padding: $spacing-lg $spacing-xl;
	margin-bottom: $spacing-xs;
	@include press-effect;

	.history-top {
		display: flex;
		justify-content: space-between;
		align-items: center;
	}
	.history-left {
		.history-date { display: block; font-size: $font-body; font-weight: 500; color: $color-label; }
		.history-sub { display: block; font-size: $font-caption2; color: $color-label-quaternary; margin-top: 2rpx; }
	}
	.history-right {
		display: flex; align-items: center; gap: $spacing-xs;
		.history-cal { font-size: $font-subhead; font-weight: 600; color: $color-accent; font-variant-numeric: tabular-nums; }
		.history-arrow { font-size: 32rpx; color: $color-separator-opaque; font-weight: 300; }
	}
	.history-macros {
		display: flex;
		align-items: center;
		gap: $spacing-xs;
		margin-top: $spacing-sm;
		padding-top: $spacing-sm;
		border-top: 0.5rpx solid $color-separator;

		.hm-item { font-size: $font-caption2; color: $color-label-quaternary; font-weight: 500; }
		.hm-sep { font-size: $font-caption2; color: $color-separator-opaque; }
	}
}

.empty {
	@include empty-state;
	.empty-icon-letter { font-size: 44rpx; font-weight: 700; color: $color-label-quaternary; }
}
</style>
