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

		<!-- 活动环卡片 - Apple Health 风格 -->
		<view class="activity-card">
			<view class="activity-header">
				<text class="activity-label">今日活动</text>
				<text class="activity-date">{{ todayStr }}</text>
			</view>
			<view class="activity-rings">
				<view class="ring-group">
					<!-- 训练环 -->
					<view class="ring-item">
						<view class="ring ring-move">
							<view class="ring-inner">
								<text class="ring-value">{{ dashboard.today_workouts || 0 }}</text>
								<text class="ring-unit">次</text>
							</view>
						</view>
						<text class="ring-label">训练</text>
					</view>
					<!-- 热量环 -->
					<view class="ring-item">
						<view class="ring ring-exercise">
							<view class="ring-inner">
								<text class="ring-value">{{ dashboard.today_calories || 0 }}</text>
								<text class="ring-unit">千卡</text>
							</view>
						</view>
						<text class="ring-label">热量</text>
					</view>
					<!-- 体重 -->
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
						<text class="quick-icon">💪</text>
					</view>
					<text class="quick-label">训练</text>
					<text class="quick-hint">记录运动</text>
				</view>
				<view class="quick-card" @tap="goTo('/pages/diet/add')">
					<view class="quick-icon-bg bg-green">
						<text class="quick-icon">🥗</text>
					</view>
					<text class="quick-label">饮食</text>
					<text class="quick-hint">营养追踪</text>
				</view>
				<view class="quick-card" @tap="goTo('/pages/weight/add')">
					<view class="quick-icon-bg bg-orange">
						<text class="quick-icon">⚖️</text>
					</view>
					<text class="quick-label">体重</text>
					<text class="quick-hint">体态管理</text>
				</view>
			</view>
		</view>

		<!-- 健康概览 -->
		<view class="section">
			<text class="section-title">健康概览</text>
			<view class="overview-list">
				<view class="overview-card" @tap="goTo('/pages/workout/list')">
					<view class="overview-left">
						<view class="overview-icon-bg bg-red">
							<text class="overview-icon">🏋️</text>
						</view>
						<view class="overview-info">
							<text class="overview-title">训练记录</text>
							<text class="overview-sub">累计 {{ dashboard.workout_count || 0 }} 次训练</text>
						</view>
					</view>
					<view class="overview-right">
						<text class="overview-arrow">
							<text class="sf-chevron">›</text>
						</text>
					</view>
				</view>
				<view class="overview-card" @tap="goTo('/pages/diet/list')">
					<view class="overview-left">
						<view class="overview-icon-bg bg-green">
							<text class="overview-icon">🍽️</text>
						</view>
						<view class="overview-info">
							<text class="overview-title">饮食记录</text>
							<text class="overview-sub">营养摄入追踪</text>
						</view>
					</view>
					<view class="overview-right">
						<text class="overview-arrow">
							<text class="sf-chevron">›</text>
						</text>
					</view>
				</view>
				<view class="overview-card" @tap="goTo('/pages/weight/list')">
					<view class="overview-left">
						<view class="overview-icon-bg bg-orange">
							<text class="overview-icon">📊</text>
						</view>
						<view class="overview-info">
							<text class="overview-title">体重趋势</text>
							<text class="overview-sub">身体变化追踪</text>
						</view>
					</view>
					<view class="overview-right">
						<text class="overview-arrow">
							<text class="sf-chevron">›</text>
						</text>
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
	import { getDashboard } from '../../api/stats'

	const userStore = useUserStore()
	const dashboard = ref({})

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

	onShow(() => loadDashboard())
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
	padding: $spacing-sm $spacing-xl $spacing-md;

	.greeting {
		.greeting-sub {
			display: block;
			font-size: $font-subhead;
			color: $color-label-quaternary;
			font-weight: 500;
		}
		.greeting-name {
			display: block;
			font-size: $font-title1;
			font-weight: 700;
			color: $color-label;
			letter-spacing: -1.5rpx;
			margin-top: 2rpx;
		}
	}

	.header-avatar {
		width: 76rpx;
		height: 76rpx;
		border-radius: 50%;
		background: linear-gradient(135deg, $color-primary, $color-indigo);
		display: flex;
		align-items: center;
		justify-content: center;
		box-shadow: 0 4rpx 20rpx rgba(0, 122, 255, 0.25);
		@include press-effect;

		.avatar-letter {
			color: #fff;
			font-size: $font-callout;
			font-weight: 600;
		}
	}
}

// --- Activity Card (Apple Health Rings Style) ---
.activity-card {
	margin: $spacing-md $spacing-xl $spacing-xl;
	@include card;
	padding: $spacing-xl;
	background: linear-gradient(145deg, #1C1C1E 0%, #2C2C2E 100%);

	.activity-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: $spacing-xl;

		.activity-label {
			font-size: $font-headline;
			font-weight: 700;
			color: #fff;
		}
		.activity-date {
			font-size: $font-footnote;
			color: rgba(255, 255, 255, 0.45);
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

		.ring-label {
			display: block;
			font-size: $font-caption1;
			color: rgba(255, 255, 255, 0.55);
			font-weight: 500;
			margin-top: $spacing-sm;
		}
	}

	.ring {
		width: 140rpx;
		height: 140rpx;
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
			border: 8rpx solid rgba(255, 255, 255, 0.08);
		}

		&::after {
			content: '';
			position: absolute;
			inset: 0;
			border-radius: 50%;
			border: 8rpx solid transparent;
		}

		.ring-inner {
			text-align: center;
		}

		.ring-value {
			display: block;
			font-size: 40rpx;
			font-weight: 700;
			color: #fff;
			letter-spacing: -1rpx;
			line-height: 1.1;
		}
		.ring-unit {
			display: block;
			font-size: $font-caption2;
			color: rgba(255, 255, 255, 0.5);
			font-weight: 500;
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
		font-size: $font-title2;
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
		padding: $spacing-lg $spacing-sm $spacing-md;
		text-align: center;
		@include press-effect;

		.quick-icon-bg {
			width: 80rpx;
			height: 80rpx;
			border-radius: $radius-lg;
			display: flex;
			align-items: center;
			justify-content: center;
			margin: 0 auto $spacing-sm;
		}

		.quick-icon { font-size: 36rpx; }

		.quick-label {
			display: block;
			font-size: $font-footnote;
			color: $color-label;
			font-weight: 600;
			margin-bottom: 2rpx;
		}
		.quick-hint {
			display: block;
			font-size: $font-caption2;
			color: $color-label-quaternary;
			font-weight: 400;
		}
	}
}

.bg-red { background: $color-red-light; }
.bg-green { background: $color-green-light; }
.bg-orange { background: $color-orange-light; }
.bg-pink { background: $color-pink-light; }

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
				width: 64rpx;
				height: 64rpx;
				border-radius: $radius-md;
				display: flex;
				align-items: center;
				justify-content: center;
				flex-shrink: 0;
			}
			.overview-icon { font-size: 30rpx; }

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
					margin-top: 2rpx;
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
</style>
