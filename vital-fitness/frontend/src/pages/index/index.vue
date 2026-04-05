<template>
	<view class="container">
		<!-- 用户信息 -->
		<view class="user-card">
			<view class="user-info">
				<text class="username">{{ userInfo.nickname || '健身达人' }}</text>
				<text class="user-stats">今日训练 {{ dashboard.today_workouts || 0 }} 次</text>
			</view>
			<view class="quick-actions">
				<button class="action-btn" @tap="goTo('/pages/workout/add')">记录训练</button>
				<button class="action-btn action-btn-green" @tap="goTo('/pages/diet/add')">记录饮食</button>
			</view>
		</view>

		<!-- 数据概览 -->
		<view class="overview">
			<view class="overview-item" @tap="goTo('/pages/workout/list')">
				<text class="item-value">{{ dashboard.workout_count || 0 }}</text>
				<text class="item-label">训练次数</text>
			</view>
			<view class="overview-item" @tap="goTo('/pages/weight/list')">
				<text class="item-value">{{ dashboard.latest_weight || '--' }}</text>
				<text class="item-label">体重(kg)</text>
			</view>
			<view class="overview-item" @tap="goTo('/pages/diet/list')">
				<text class="item-value">{{ dashboard.today_calories || 0 }}</text>
				<text class="item-label">今日热量</text>
			</view>
		</view>

		<!-- 快捷入口 -->
		<view class="shortcuts">
			<view class="shortcut-item" @tap="goTo('/pages/weight/add')">
				<text class="shortcut-icon">⚖️</text>
				<text class="shortcut-text">记录体重</text>
			</view>
			<view class="shortcut-item" @tap="goTo('/pages/mood/add')">
				<text class="shortcut-icon">😊</text>
				<text class="shortcut-text">记录心情</text>
			</view>
			<view class="shortcut-item" @tap="goTo('/pages/statistics/index')">
				<text class="shortcut-icon">📊</text>
				<text class="shortcut-text">数据统计</text>
			</view>
			<view class="shortcut-item" @tap="goTo('/pages/user/profile')">
				<text class="shortcut-icon">👤</text>
				<text class="shortcut-text">个人资料</text>
			</view>
		</view>
	</view>
</template>

<script>
	import { getDashboard } from '../../api/stats'

	export default {
		data() {
			return { dashboard: {}, userInfo: {} }
		},
		onShow() {
			this.userInfo = this.$store.getters.userInfo || {}
			this.loadDashboard()
		},
		methods: {
			async loadDashboard() {
				try {
					const res = await getDashboard()
					this.dashboard = res.data || {}
				} catch (e) {}
			},
			goTo(url) { uni.navigateTo({ url }) }
		}
	}
</script>

<style lang="scss" scoped>
.container { padding: 20rpx; }

.user-card {
	background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
	border-radius: 20rpx; padding: 40rpx; margin-bottom: 20rpx; color: white;
	.user-info { margin-bottom: 30rpx;
		.username { font-size: 36rpx; font-weight: bold; display: block; }
		.user-stats { font-size: 24rpx; opacity: 0.9; margin-top: 10rpx; display: block; }
	}
	.quick-actions { display: flex; gap: 20rpx;
		.action-btn { flex: 1; background: rgba(255,255,255,0.2); color: white; border: none; border-radius: 15rpx; font-size: 28rpx; padding: 15rpx 0; }
		.action-btn-green { background: rgba(107,207,127,0.4); }
	}
}

.overview {
	display: flex; gap: 15rpx; margin-bottom: 20rpx;
	.overview-item {
		flex: 1; background: white; border-radius: 20rpx; padding: 30rpx; text-align: center;
		.item-value { font-size: 40rpx; font-weight: bold; color: #333; display: block; }
		.item-label { font-size: 22rpx; color: #999; margin-top: 8rpx; display: block; }
	}
}

.shortcuts {
	display: flex; flex-wrap: wrap; gap: 15rpx;
	.shortcut-item {
		width: calc(50% - 8rpx); background: white; border-radius: 20rpx; padding: 30rpx; text-align: center;
		.shortcut-icon { font-size: 48rpx; display: block; margin-bottom: 10rpx; }
		.shortcut-text { font-size: 26rpx; color: #333; }
	}
}
</style>
