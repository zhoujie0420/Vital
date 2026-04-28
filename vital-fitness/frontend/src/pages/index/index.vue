<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<!-- Header -->
		<view class="header">
			<view class="greeting">
				<text class="greeting-sub">{{ greetingTime }}</text>
				<text class="greeting-name">{{ userStore.userInfo?.nickname || '健身达人' }}</text>
			</view>
			<view class="header-avatar" @tap="goTo('/pages/user/index')">
				<text class="avatar-letter">{{ (userStore.userInfo?.nickname || '你')[0] }}</text>
			</view>
		</view>

		<!-- 活动环卡片 -->
		<view class="activity-card">
			<view class="activity-header">
				<text class="activity-label">今日活动</text>
				<text class="activity-date">{{ todayStr }}</text>
			</view>
			<view class="activity-rings">
				<view class="ring-group">
					<view class="ring-item">
						<view class="ring ring-move">
							<view class="ring-inner">
								<text class="ring-value">{{ dashboard.today_workouts || 0 }}</text>
								<text class="ring-unit">次</text>
							</view>
						</view>
						<text class="ring-label">训练</text>
					</view>
					<view class="ring-item">
						<view class="ring ring-exercise">
							<view class="ring-inner">
								<text class="ring-value">{{ dashboard.today_calories || 0 }}</text>
								<text class="ring-unit">千卡</text>
							</view>
						</view>
						<text class="ring-label">热量</text>
					</view>
					<view class="ring-item">
						<view class="ring ring-stand">
							<view class="ring-inner">
								<text class="ring-value">{{ dashboard.latest_weight || '--' }}</text>
								<text class="ring-unit">kg</text>
							</view>
						</view>
						<text class="ring-label">体重</text>
					</view>
				</view>
			</view>
		</view>

		<!-- 快捷操作 -->
		<view class="section">
			<text class="section-title">快捷记录</text>
			<view class="quick-grid">
				<view class="quick-card" @tap="goTo('/pages/workout/add')">
					<view class="quick-icon-bg bg-red">
						<text class="quick-icon iconfont">&#xe6a3;</text>
					</view>
					<text class="quick-label">训练</text>
					<text class="quick-hint">记录运动</text>
				</view>
				<view class="quick-card" @tap="goTo('/pages/diet/add')">
					<view class="quick-icon-bg bg-green">
						<text class="quick-icon iconfont">&#xe6a1;</text>
					</view>
					<text class="quick-label">饮食</text>
					<text class="quick-hint">营养追踪</text>
				</view>
				<view class="quick-card" @tap="goTo('/pages/weight/add')">
					<view class="quick-icon-bg bg-orange">
						<text class="quick-icon iconfont">&#xe6a2;</text>
					</view>
					<text class="quick-label">体重</text>
					<text class="quick-hint">体态管理</text>
				</view>
			</view>
		</view>

		<!-- 饮食规划模块 -->
		<view class="section" v-if="dietPlanStore.activePlan">
			<view class="section-header">
				<text class="section-title">今日饮食</text>
				<text class="section-link" @tap="goTo('/pages/diet/plan')">方案设置 ›</text>
			</view>
			<view class="diet-plan-card">
				<view class="dp-remaining">
					<text class="dp-remaining-val" :class="{ 'dp-over': remainingCal < 0 }">{{ Math.abs(remainingCal) }}</text>
					<text class="dp-remaining-label">{{ remainingCal >= 0 ? '剩余千卡' : '超出千卡' }}</text>
				</view>
				<view class="dp-macros">
					<view class="dp-macro">
						<view class="dp-macro-track"><view class="dp-macro-fill fill-teal" :style="{ width: macroPct(todayConsumed.protein, todayTargets.protein) }"></view></view>
						<text class="dp-macro-text">蛋白 {{ todayConsumed.protein }}/{{ todayTargets.protein }}g</text>
					</view>
					<view class="dp-macro">
						<view class="dp-macro-track"><view class="dp-macro-fill fill-orange" :style="{ width: macroPct(todayConsumed.carbs, todayTargets.carbs) }"></view></view>
						<text class="dp-macro-text">碳水 {{ todayConsumed.carbs }}/{{ todayTargets.carbs }}g</text>
					</view>
					<view class="dp-macro">
						<view class="dp-macro-track"><view class="dp-macro-fill fill-yellow" :style="{ width: macroPct(todayConsumed.fat, todayTargets.fat) }"></view></view>
						<text class="dp-macro-text">脂肪 {{ todayConsumed.fat }}/{{ todayTargets.fat }}g</text>
					</view>
				</view>
				<view class="dp-meals">
					<view class="dp-meal" @tap="goTo('/pages/diet/add?meal=breakfast')"><text>早餐</text></view>
					<view class="dp-meal" @tap="goTo('/pages/diet/add?meal=lunch')"><text>午餐</text></view>
					<view class="dp-meal" @tap="goTo('/pages/diet/add?meal=dinner')"><text>晚餐</text></view>
					<view class="dp-meal" @tap="goTo('/pages/diet/add?meal=snack')"><text>加餐</text></view>
				</view>
			</view>
		</view>

		<view class="section" v-else-if="!dietPlanStore.activePlan">
			<view class="diet-plan-empty" @tap="goTo('/pages/diet/plan')">
				<text class="dpe-text">设置饮食规划，追踪每日营养目标</text>
				<text class="dpe-link">去设置 ›</text>
			</view>
		</view>

		<!-- 健康概览 -->
		<view class="section">
			<text class="section-title">健康概览</text>
			<view class="overview-list">
				<view class="overview-card" @tap="goTo('/pages/workout/list')">
					<view class="overview-left">
						<view class="overview-icon-bg bg-red">
							<text class="overview-icon-text">W</text>
						</view>
						<view class="overview-info">
							<text class="overview-title">训练记录</text>
							<text class="overview-sub">累计 {{ dashboard.workout_count || 0 }} 次训练</text>
						</view>
					</view>
					<view class="overview-right">
						<text class="sf-chevron">›</text>
					</view>
				</view>
				<view class="overview-card" @tap="goTo('/pages/diet/list')">
					<view class="overview-left">
						<view class="overview-icon-bg bg-green">
							<text class="overview-icon-text">D</text>
						</view>
						<view class="overview-info">
							<text class="overview-title">饮食记录</text>
							<text class="overview-sub">营养摄入追踪</text>
						</view>
					</view>
					<view class="overview-right">
						<text class="sf-chevron">›</text>
					</view>
				</view>
				<view class="overview-card" @tap="goTo('/pages/weight/list')">
					<view class="overview-left">
						<view class="overview-icon-bg bg-orange">
							<text class="overview-icon-text">B</text>
						</view>
						<view class="overview-info">
							<text class="overview-title">体重趋势</text>
							<text class="overview-sub">身体变化追踪</text>
						</view>
					</view>
					<view class="overview-right">
						<text class="sf-chevron">›</text>
					</view>
				</view>
			</view>
		</view>
	</view>
</template>

<script setup>
	import { ref, computed } from 'vue'
	import { onShow } from '@dcloudio/uni-app'
	import { useUserStore } from '../../store'
	import { useDietPlanStore } from '../../store/dietPlan'
	import { getDashboard } from '../../api/stats'
	import { getDietRecords } from '../../api/diet'
	import { getTodayTargets, calculateRemaining } from '../../utils/dietPlanCalc'

	const userStore = useUserStore()
	const dietPlanStore = useDietPlanStore()
	const dashboard = ref({})
	const todayDietRecords = ref([])

	const todayTargets = computed(() => {
		if (!dietPlanStore.activePlan) return { calories: 0, protein: 0, carbs: 0, fat: 0 }
		return getTodayTargets(dietPlanStore.activePlan, new Date())
	})

	const todayConsumed = computed(() => {
		const records = todayDietRecords.value
		return {
			calories: Math.round(records.reduce((s, r) => s + (r.total_calories || 0), 0)),
			protein: Math.round(records.reduce((s, r) => s + (r.protein || 0), 0)),
			carbs: Math.round(records.reduce((s, r) => s + (r.carbs || 0), 0)),
			fat: Math.round(records.reduce((s, r) => s + (r.fat || 0), 0))
		}
	})

	const remainingCal = computed(() => {
		return Math.round(todayTargets.value.calories - todayConsumed.value.calories)
	})

	const macroPct = (consumed, target) => {
		if (!target || target <= 0) return '0%'
		return Math.min(Math.round((consumed / target) * 100), 100) + '%'
	}

	const topPadding = computed(() => {
		const app = getApp()
		return (app.globalData?.customBarHeight || 88) + 8
	})

	const greetingTime = computed(() => {
		const h = new Date().getHours()
		if (h < 6) return '夜深了'
		if (h < 12) return '早上好'
		if (h < 18) return '下午好'
		return '晚上好'
	})

	const todayStr = computed(() => {
		const d = new Date()
		return `${d.getMonth() + 1}月${d.getDate()}日`
	})

	const goTo = (url) => {
		const tabbarPages = ['/pages/index/index', '/pages/workout/list', '/pages/diet/list', '/pages/statistics/index', '/pages/user/index']
		if (tabbarPages.includes(url)) {
			uni.switchTab({ url })
		} else {
			uni.navigateTo({ url })
		}
	}

	const loadDashboard = async () => {
		try {
			const res = await getDashboard()
			dashboard.value = res.data || {}
		} catch (e) {}
	}

	onShow(() => {
		loadDashboard()
		dietPlanStore.init()
		loadTodayDiet()
	})

	const loadTodayDiet = async () => {
		try {
			const res = await getDietRecords({ page: 1, page_size: 50 })
			const all = res.data || []
			const todayKey = new Date().toISOString().split('T')[0]
			todayDietRecords.value = all.filter(r => {
				const d = new Date(r.record_date)
				return d.toISOString().split('T')[0] === todayKey
			})
		} catch (e) {}
	}
</script>

<style lang="scss" scoped>
@import '../../styles/variables.scss';

.page {
	min-height: 100vh;
	background: $color-bg;
	padding-bottom: 40rpx;
}

// --- Header ---
.header {
	display: flex;
	justify-content: space-between;
	align-items: center;
	padding: $spacing-sm $spacing-xl $spacing-lg;

	.greeting {
		.greeting-sub {
			display: block;
			font-size: $font-footnote;
			color: $color-label-quaternary;
			font-weight: 500;
			letter-spacing: 0.5rpx;
		}
		.greeting-name {
			display: block;
			font-size: $font-title1;
			font-weight: 700;
			color: $color-label;
			letter-spacing: -1.5rpx;
			margin-top: 4rpx;
		}
	}

	.header-avatar {
		width: 80rpx;
		height: 80rpx;
		border-radius: 50%;
		background: $color-primary;
		display: flex;
		align-items: center;
		justify-content: center;
		box-shadow: 0 4rpx 16rpx rgba(16, 185, 129, 0.25);
		@include press-effect;

		.avatar-letter {
			color: #fff;
			font-size: $font-headline;
			font-weight: 600;
		}
	}
}

// --- Activity Card ---
.activity-card {
	margin: $spacing-sm $spacing-xl $spacing-xl;
	border-radius: $radius-2xl;
	padding: $spacing-xl $spacing-xl $spacing-2xl;
	background: $color-label;
	box-shadow: 0 8rpx 32rpx rgba(24, 24, 27, 0.15);

	.activity-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: $spacing-xl;

		.activity-label {
			font-size: $font-headline;
			font-weight: 700;
			color: #fff;
			letter-spacing: -0.5rpx;
		}
		.activity-date {
			font-size: $font-caption1;
			color: rgba(255, 255, 255, 0.4);
			font-weight: 500;
		}
	}

	.ring-group {
		display: flex;
		justify-content: space-around;
		align-items: center;
	}

	.ring-item {
		text-align: center;
		@include stagger-fade(3);

		.ring-label {
			display: block;
			font-size: $font-caption1;
			color: rgba(255, 255, 255, 0.45);
			font-weight: 500;
			margin-top: $spacing-sm;
			letter-spacing: 0.5rpx;
		}
	}

	.ring {
		width: 148rpx;
		height: 148rpx;
		border-radius: 50%;
		display: flex;
		align-items: center;
		justify-content: center;
		position: relative;

		&::before {
			content: '';
			position: absolute;
			inset: 0;
			border-radius: 50%;
			border: 6rpx solid rgba(255, 255, 255, 0.06);
		}

		&::after {
			content: '';
			position: absolute;
			inset: 0;
			border-radius: 50%;
			border: 6rpx solid transparent;
		}

		.ring-inner {
			text-align: center;
		}

		.ring-value {
			display: block;
			font-size: 42rpx;
			font-weight: 700;
			color: #fff;
			letter-spacing: -1rpx;
			line-height: 1.1;
			font-variant-numeric: tabular-nums;
		}
		.ring-unit {
			display: block;
			font-size: $font-caption2;
			color: rgba(255, 255, 255, 0.4);
			font-weight: 500;
			margin-top: 2rpx;
		}
	}

	.ring-move {
		&::after { border-color: $color-red; border-bottom-color: transparent; }
	}
	.ring-exercise {
		&::after { border-color: $color-green; border-bottom-color: transparent; }
	}
	.ring-stand {
		&::after { border-color: $color-teal; border-bottom-color: transparent; }
	}
}

// --- Quick Actions ---
.section {
	padding: 0 $spacing-xl;
	margin-bottom: $spacing-xl;

	.section-title {
		display: block;
		font-size: $font-title3;
		font-weight: 700;
		color: $color-label;
		margin-bottom: $spacing-md;
		letter-spacing: -0.5rpx;
	}
}

.quick-grid {
	display: flex;
	gap: $spacing-sm;

	.quick-card {
		flex: 1;
		@include card;
		padding: $spacing-xl $spacing-sm $spacing-lg;
		text-align: center;
		@include press-effect;
		@include stagger-fade(3);

		.quick-icon-bg {
			width: 88rpx;
			height: 88rpx;
			border-radius: $radius-xl;
			display: flex;
			align-items: center;
			justify-content: center;
			margin: 0 auto $spacing-sm;
		}

		.quick-icon {
			font-size: 40rpx;
			font-style: normal;
		}

		.quick-label {
			display: block;
			font-size: $font-subhead;
			color: $color-label;
			font-weight: 600;
			margin-bottom: 4rpx;
		}
		.quick-hint {
			display: block;
			font-size: $font-caption2;
			color: $color-label-quaternary;
			font-weight: 400;
		}
	}
}

.bg-red {
	background: $color-red-light;
	.quick-icon, .overview-icon-text { color: $color-red; }
}
.bg-green {
	background: $color-green-light;
	.quick-icon, .overview-icon-text { color: $color-green; }
}
.bg-orange {
	background: $color-accent-light;
	.quick-icon, .overview-icon-text { color: $color-accent; }
}

// --- Overview List ---
.overview-list {
	@include card;
	padding: 0;
	overflow: hidden;

	.overview-card {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: $spacing-lg $spacing-xl;
		position: relative;
		@include press-effect;

		&::after {
			content: '';
			position: absolute;
			bottom: 0;
			left: 108rpx;
			right: $spacing-xl;
			height: 0.5rpx;
			background: $color-separator;
		}

		&:last-child::after { display: none; }

		.overview-left {
			display: flex;
			align-items: center;
			gap: $spacing-md;

			.overview-icon-bg {
				width: 68rpx;
				height: 68rpx;
				border-radius: $radius-md;
				display: flex;
				align-items: center;
				justify-content: center;
				flex-shrink: 0;
			}
			.overview-icon-text {
				font-size: $font-headline;
				font-weight: 700;
			}

			.overview-info {
				.overview-title {
					display: block;
					font-size: $font-body;
					font-weight: 500;
					color: $color-label;
				}
				.overview-sub {
					display: block;
					font-size: $font-caption1;
					color: $color-label-quaternary;
					margin-top: 4rpx;
				}
			}
		}

		.overview-right {
			.sf-chevron {
				font-size: 36rpx;
				color: $color-separator-opaque;
				font-weight: 300;
			}
		}
	}
}

// --- Diet Plan Module ---
.section-header {
	display: flex;
	justify-content: space-between;
	align-items: baseline;
	margin-bottom: $spacing-md;

	.section-link {
		font-size: $font-caption1;
		color: $color-primary;
		font-weight: 500;
	}
}

.diet-plan-card {
	@include card;
	padding: $spacing-xl;

	.dp-remaining {
		text-align: center;
		margin-bottom: $spacing-lg;

		.dp-remaining-val {
			display: block;
			font-size: 64rpx;
			font-weight: 700;
			color: $color-primary;
			letter-spacing: -2rpx;
			font-variant-numeric: tabular-nums;
			line-height: 1.1;
		}
		.dp-over { color: $color-red; }
		.dp-remaining-label {
			display: block;
			font-size: $font-caption1;
			color: $color-label-quaternary;
			margin-top: $spacing-xs;
		}
	}

	.dp-macros {
		display: flex;
		flex-direction: column;
		gap: $spacing-sm;
		margin-bottom: $spacing-lg;
	}

	.dp-macro {
		.dp-macro-track {
			height: 12rpx;
			background: $color-fill;
			border-radius: 6rpx;
			overflow: hidden;
			margin-bottom: 4rpx;
		}
		.dp-macro-fill {
			height: 100%;
			border-radius: 6rpx;
			min-width: 4rpx;
			transition: width 0.4s cubic-bezier(0.16, 1, 0.3, 1);
		}
		.fill-teal { background: $color-teal; }
		.fill-orange { background: $color-orange; }
		.fill-yellow { background: $color-yellow; }
		.dp-macro-text {
			font-size: $font-caption2;
			color: $color-label-tertiary;
			font-weight: 500;
			font-variant-numeric: tabular-nums;
		}
	}

	.dp-meals {
		display: flex;
		gap: $spacing-sm;

		.dp-meal {
			flex: 1;
			text-align: center;
			padding: $spacing-md 0;
			background: $color-fill;
			border-radius: $radius-md;
			@include press-effect;

			text {
				font-size: $font-caption1;
				color: $color-label-tertiary;
				font-weight: 500;
			}
		}
	}
}

.diet-plan-empty {
	@include card;
	padding: $spacing-xl;
	text-align: center;
	@include press-effect;

	.dpe-text {
		display: block;
		font-size: $font-subhead;
		color: $color-label-quaternary;
		margin-bottom: $spacing-sm;
	}
	.dpe-link {
		font-size: $font-subhead;
		color: $color-primary;
		font-weight: 600;
	}
}
</style>
