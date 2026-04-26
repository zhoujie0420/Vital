<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="header">
			<view class="greeting">
				<text class="greeting-sub">{{ greetingTime }}</text>
				<text class="greeting-name">{{ userStore.userInfo?.nickname || '健身达人' }}</text>
			</view>
			<view class="header-avatar" @tap="goTo('/pages/user/index')">
				<text class="avatar-letter">{{ (userStore.userInfo?.nickname || '你')[0] }}</text>
			</view>
		</view>

		<!-- 今日概览卡片 -->
		<view class="hero-card">
			<view class="hero-bg"></view>
			<view class="hero-content">
				<text class="hero-label">今日概览</text>
				<view class="hero-stats">
					<view class="hero-stat">
						<view class="hero-ring">
							<text class="hero-ring-val">{{ dashboard.today_workouts || 0 }}</text>
						</view>
						<text class="hero-stat-label">训练</text>
					</view>
					<view class="hero-stat">
						<view class="hero-ring ring-orange">
							<text class="hero-ring-val">{{ dashboard.today_calories || 0 }}</text>
						</view>
						<text class="hero-stat-label">千卡</text>
					</view>
					<view class="hero-stat">
						<view class="hero-ring ring-green">
							<text class="hero-ring-val">{{ dashboard.latest_weight || '--' }}</text>
						</view>
						<text class="hero-stat-label">体重kg</text>
					</view>
				</view>
			</view>
		</view>

		<!-- 快捷操作 -->
		<view class="section">
			<text class="section-title">快捷记录</text>
			<view class="action-grid">
				<view class="action-card action-workout" @tap="goTo('/pages/workout/add')">
					<view class="action-icon-wrap"><text class="action-icon">💪</text></view>
					<text class="action-label">训练</text>
				</view>
				<view class="action-card action-diet" @tap="goTo('/pages/diet/add')">
					<view class="action-icon-wrap"><text class="action-icon">🥗</text></view>
					<text class="action-label">饮食</text>
				</view>
				<view class="action-card action-weight" @tap="goTo('/pages/weight/add')">
					<view class="action-icon-wrap"><text class="action-icon">⚖️</text></view>
					<text class="action-label">体重</text>
				</view>
				<view class="action-card action-mood" @tap="goTo('/pages/mood/add')">
					<view class="action-icon-wrap"><text class="action-icon">😊</text></view>
					<text class="action-label">心情</text>
				</view>
			</view>
		</view>

		<!-- 数据卡片 -->
		<view class="section">
			<text class="section-title">数据概览</text>
			<view class="data-cards">
				<view class="data-card" @tap="goTo('/pages/workout/list')">
					<view class="data-card-left">
						<text class="data-card-icon">🏋️</text>
						<view>
							<text class="data-card-title">训练次数</text>
							<text class="data-card-sub">累计记录</text>
						</view>
					</view>
					<view class="data-card-right">
						<text class="data-card-value">{{ dashboard.workout_count || 0 }}</text>
						<text class="data-card-arrow">›</text>
					</view>
				</view>
				<view class="data-card" @tap="goTo('/pages/diet/list')">
					<view class="data-card-left">
						<text class="data-card-icon">🍽️</text>
						<view>
							<text class="data-card-title">饮食记录</text>
							<text class="data-card-sub">营养追踪</text>
						</view>
					</view>
					<view class="data-card-right">
						<text class="data-card-arrow">›</text>
					</view>
				</view>
				<view class="data-card" @tap="goTo('/pages/weight/list')">
					<view class="data-card-left">
						<text class="data-card-icon">📊</text>
						<view>
							<text class="data-card-title">体重趋势</text>
							<text class="data-card-sub">变化追踪</text>
						</view>
					</view>
					<view class="data-card-right">
						<text class="data-card-arrow">›</text>
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
.page {
	min-height: 100vh;
	background: linear-gradient(180deg, #e8f3ff 0%, #f2f2f7 320rpx);
	padding-bottom: 40rpx;
}

.header {
	display: flex;
	justify-content: space-between;
	align-items: center;
	padding: 12rpx 32rpx 16rpx;

	.greeting {
		.greeting-sub {
			display: block;
			font-size: 28rpx;
			color: #8e8e93;
			font-weight: 500;
		}
		.greeting-name {
			display: block;
			font-size: 52rpx;
			font-weight: 700;
			color: #1c1c1e;
			letter-spacing: -1rpx;
			margin-top: 4rpx;
		}
	}

	.header-avatar {
		width: 80rpx;
		height: 80rpx;
		border-radius: 50%;
		background: linear-gradient(135deg, #007aff, #5856d6);
		display: flex;
		align-items: center;
		justify-content: center;
		box-shadow: 0 4rpx 16rpx rgba(0, 122, 255, 0.3);

		&:active { transform: scale(0.92); }

		.avatar-letter {
			color: #fff;
			font-size: 32rpx;
			font-weight: 600;
		}
	}
}

.hero-card {
	margin: 16rpx 32rpx 32rpx;
	border-radius: 28rpx;
	overflow: hidden;
	position: relative;
	box-shadow: 0 12rpx 40rpx rgba(0, 122, 255, 0.35);

	.hero-bg {
		position: absolute;
		top: 0; left: 0; right: 0; bottom: 0;
		background: linear-gradient(135deg, #0a84ff 0%, #5856d6 50%, #af52de 100%);
	}

	.hero-content {
		position: relative;
		z-index: 1;
		padding: 40rpx 36rpx;

		.hero-label {
			display: block;
			font-size: 24rpx;
			color: rgba(255, 255, 255, 0.7);
			font-weight: 600;
			letter-spacing: 2rpx;
			text-transform: uppercase;
			margin-bottom: 32rpx;
		}

		.hero-stats {
			display: flex;
			justify-content: space-around;
		}

		.hero-stat {
			text-align: center;

			.hero-ring {
				width: 128rpx;
				height: 128rpx;
				border-radius: 50%;
				background: rgba(255, 255, 255, 0.15);
				border: 3rpx solid rgba(255, 255, 255, 0.35);
				display: flex;
				align-items: center;
				justify-content: center;
				margin: 0 auto 16rpx;
				box-shadow: 0 4rpx 16rpx rgba(0, 0, 0, 0.15);

				.hero-ring-val {
					font-size: 40rpx;
					font-weight: 800;
					color: #fff;
					letter-spacing: -1rpx;
				}
			}

			.ring-orange {
				border-color: rgba(255, 190, 40, 0.55);
				background: rgba(255, 159, 10, 0.18);
			}

			.ring-green {
				border-color: rgba(52, 199, 89, 0.55);
				background: rgba(48, 209, 88, 0.18);
			}

			.hero-stat-label {
				font-size: 24rpx;
				color: rgba(255, 255, 255, 0.85);
				font-weight: 600;
				letter-spacing: 0.5rpx;
			}
		}
	}
}

.section {
	padding: 0 32rpx;
	margin-bottom: 32rpx;

	.section-title {
		display: block;
		font-size: 36rpx;
		font-weight: 700;
		color: #1c1c1e;
		margin-bottom: 20rpx;
		letter-spacing: -0.5rpx;
	}
}

.action-grid {
	display: flex;
	gap: 16rpx;

	.action-card {
		flex: 1;
		border-radius: 22rpx;
		padding: 28rpx 0 22rpx;
		text-align: center;
		transition: transform 0.15s ease;

		&:active { transform: scale(0.94); }

		.action-icon-wrap {
			width: 72rpx;
			height: 72rpx;
			border-radius: 20rpx;
			background: rgba(255, 255, 255, 0.22);
			display: flex;
			align-items: center;
			justify-content: center;
			margin: 0 auto 14rpx;
		}

		.action-icon { font-size: 36rpx; }
		.action-label { display: block; font-size: 24rpx; color: #fff; font-weight: 700; }
	}

	.action-workout {
		background: linear-gradient(145deg, #0a84ff, #005ec7);
		box-shadow: 0 8rpx 24rpx rgba(0, 122, 255, 0.38);
	}
	.action-diet {
		background: linear-gradient(145deg, #30d158, #1ea33e);
		box-shadow: 0 8rpx 24rpx rgba(52, 199, 89, 0.38);
	}
	.action-weight {
		background: linear-gradient(145deg, #ff9f0a, #d97400);
		box-shadow: 0 8rpx 24rpx rgba(255, 149, 0, 0.38);
	}
	.action-mood {
		background: linear-gradient(145deg, #ff375f, #c4003d);
		box-shadow: 0 8rpx 24rpx rgba(255, 45, 85, 0.38);
	}
}

.data-cards {
	.data-card {
		display: flex;
		justify-content: space-between;
		align-items: center;
		background: #fff;
		border-radius: 16rpx;
		padding: 28rpx 28rpx 28rpx 24rpx;
		margin-bottom: 12rpx;
		box-shadow: 0 2rpx 12rpx rgba(0, 0, 0, 0.05);
		transition: transform 0.15s ease;
		border-left: 6rpx solid transparent;

		&:active { transform: scale(0.98); }

		&:nth-child(1) { border-left-color: #0a84ff; }
		&:nth-child(2) { border-left-color: #30d158; }
		&:nth-child(3) { border-left-color: #ff9f0a; }

		.data-card-left {
			display: flex;
			align-items: center;
			gap: 20rpx;

			.data-card-icon { font-size: 40rpx; }
			.data-card-title { display: block; font-size: 30rpx; font-weight: 600; color: #1c1c1e; }
			.data-card-sub { display: block; font-size: 22rpx; color: #8e8e93; margin-top: 2rpx; }
		}

		.data-card-right {
			display: flex;
			align-items: center;
			gap: 12rpx;

			.data-card-value { font-size: 40rpx; font-weight: 700; color: #007aff; }
			.data-card-arrow { font-size: 36rpx; color: #c7c7cc; }
		}
	}
}
</style>
