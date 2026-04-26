<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="banner-bg"></view>
		<view class="page-header">
			<text class="page-title">我的</text>
		</view>

		<view class="user-card" @tap="goTo('/pages/user/profile')">
			<view class="avatar">
				<text class="avatar-text">{{ (userStore.userInfo?.nickname || '用')[0] }}</text>
			</view>
			<view class="user-info">
				<text class="user-name">{{ userStore.userInfo?.nickname || '微信用户' }}</text>
				<text class="user-sub">查看和编辑个人资料</text>
			</view>
			<text class="arrow">›</text>
		</view>

		<view class="menu-group">
			<view class="menu-item" @tap="goTo('/pages/user/profile')">
				<view class="menu-left">
					<view class="menu-icon icon-blue"><text>👤</text></view>
					<text class="menu-text">个人资料</text>
				</view>
				<text class="arrow">›</text>
			</view>
			<view class="menu-item" @tap="goTo('/pages/statistics/index')">
				<view class="menu-left">
					<view class="menu-icon icon-green"><text>📊</text></view>
					<text class="menu-text">数据统计</text>
				</view>
				<text class="arrow">›</text>
			</view>
			<view class="menu-item last" @tap="goTo('/pages/user/settings')">
				<view class="menu-left">
					<view class="menu-icon icon-gray"><text>⚙️</text></view>
					<text class="menu-text">设置</text>
				</view>
				<text class="arrow">›</text>
			</view>
		</view>

		<view class="logout-wrap">
			<view class="logout-btn" @tap="logout">
				<text>退出登录</text>
			</view>
		</view>
	</view>
</template>

<script setup>
	import { computed } from 'vue'
	import { useUserStore } from '../../store'

	const userStore = useUserStore()

	const topPadding = computed(() => {
		const app = getApp()
		return (app.globalData?.customBarHeight || 88) + 8
	})

	const goTo = (url) => {
		const tabbarPages = ['/pages/index/index', '/pages/workout/list', '/pages/diet/list', '/pages/statistics/index', '/pages/user/index']
		if (tabbarPages.includes(url)) {
			uni.switchTab({ url })
		} else {
			uni.navigateTo({ url })
		}
	}

	const logout = () => {
		uni.showModal({
			title: '确认退出',
			content: '确定要退出登录吗？',
			success: (res) => {
				if (res.confirm) {
					userStore.logout()
					uni.redirectTo({ url: '/pages/auth/login' })
				}
			}
		})
	}
</script>

<style lang="scss" scoped>
.page {
	padding: 0 32rpx;
	padding-bottom: 40rpx;
	min-height: 100vh;
	background: #f2f2f7;
	position: relative;
}

.banner-bg {
	position: absolute;
	top: 0;
	left: 0;
	right: 0;
	height: 380rpx;
	background: linear-gradient(135deg, #007aff 0%, #5856d6 60%, #af52de 100%);
	z-index: 0;
}

.page-header {
	position: relative;
	z-index: 1;
	margin-bottom: 28rpx;

	.page-title {
		font-size: 52rpx;
		font-weight: 700;
		color: #fff;
		letter-spacing: -1rpx;
	}
}

.user-card {
	display: flex;
	align-items: center;
	background: rgba(255, 255, 255, 0.92);
	backdrop-filter: blur(20px);
	-webkit-backdrop-filter: blur(20px);
	border-radius: 20rpx;
	padding: 32rpx;
	margin-bottom: 24rpx;
	box-shadow: 0 8rpx 32rpx rgba(0, 0, 0, 0.12);
	transition: transform 0.15s ease;
	position: relative;
	z-index: 1;

	&:active { transform: scale(0.98); }

	.avatar {
		width: 100rpx;
		height: 100rpx;
		border-radius: 50%;
		background: linear-gradient(135deg, #007aff, #5856d6);
		display: flex;
		align-items: center;
		justify-content: center;
		margin-right: 24rpx;
		box-shadow: 0 4rpx 16rpx rgba(0, 122, 255, 0.3);

		.avatar-text {
			color: #fff;
			font-size: 40rpx;
			font-weight: 600;
		}
	}

	.user-info {
		flex: 1;

		.user-name {
			display: block;
			font-size: 36rpx;
			font-weight: 600;
			color: #1c1c1e;
		}
		.user-sub {
			display: block;
			font-size: 24rpx;
			color: #8e8e93;
			margin-top: 4rpx;
		}
	}

	.arrow {
		font-size: 40rpx;
		color: #c7c7cc;
		font-weight: 300;
	}
}

.menu-group {
	background: #fff;
	border-radius: 20rpx;
	margin-bottom: 24rpx;
	overflow: hidden;
	box-shadow: 0 2rpx 16rpx rgba(0, 0, 0, 0.04);

	.menu-item {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 24rpx 32rpx;
		border-bottom: 1rpx solid #f2f2f7;
		transition: background 0.15s ease;

		&:active { background: #f8f8f8; }
		&.last { border-bottom: none; }

		.menu-left {
			display: flex;
			align-items: center;
			gap: 20rpx;

			.menu-icon {
				width: 56rpx;
				height: 56rpx;
				border-radius: 14rpx;
				display: flex;
				align-items: center;
				justify-content: center;
				font-size: 28rpx;
			}
			.icon-blue { background: rgba(0, 122, 255, 0.1); }
			.icon-green { background: rgba(52, 199, 89, 0.1); }
			.icon-gray { background: rgba(142, 142, 147, 0.1); }

			.menu-text {
				font-size: 30rpx;
				color: #1c1c1e;
				font-weight: 500;
			}
		}

		.arrow {
			font-size: 36rpx;
			color: #c7c7cc;
			font-weight: 300;
		}
	}
}

.logout-wrap {
	padding: 0 32rpx;

	.logout-btn {
		background: #fff;
		border-radius: 16rpx;
		padding: 28rpx 0;
		text-align: center;
		box-shadow: 0 2rpx 16rpx rgba(0, 0, 0, 0.04);
		transition: transform 0.15s ease;

		&:active { transform: scale(0.98); }

		text {
			font-size: 30rpx;
			color: #ff3b30;
			font-weight: 500;
		}
	}
}
</style>
