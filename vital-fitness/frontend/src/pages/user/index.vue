<template>
	<view class="container">
		<view class="user-card" @tap="goTo('/pages/user/profile')">
			<view class="avatar">
				<text class="avatar-text">{{ (userStore.userInfo?.nickname || '用')[0] }}</text>
			</view>
			<view class="user-info">
				<text class="user-name">{{ userStore.userInfo?.nickname || '微信用户' }}</text>
			</view>
			<text class="arrow">›</text>
		</view>

		<view class="menu-group">
			<view class="menu-item" @tap="goTo('/pages/user/profile')">
				<text>个人资料</text><text class="arrow">›</text>
			</view>
			<view class="menu-item" @tap="goTo('/pages/statistics/index')">
				<text>数据统计</text><text class="arrow">›</text>
			</view>
			<view class="menu-item" @tap="goTo('/pages/user/settings')">
				<text>设置</text><text class="arrow">›</text>
			</view>
		</view>

		<button class="logout-btn" @tap="logout">退出登录</button>
	</view>
</template>

<script setup>
	import { useUserStore } from '../../store'

	const userStore = useUserStore()

	const goTo = (url) => uni.navigateTo({ url })

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
.container { padding: 32rpx; }
.user-card {
	display: flex; align-items: center; background: #fff; border-radius: 20rpx;
	padding: 32rpx; margin-bottom: 24rpx;
	.avatar { width: 96rpx; height: 96rpx; border-radius: 50%; background: linear-gradient(135deg, #007aff, #5856d6);
		display: flex; align-items: center; justify-content: center; margin-right: 24rpx;
		.avatar-text { color: #fff; font-size: 36rpx; font-weight: 600; }
	}
	.user-info { flex: 1;
		.user-name { font-size: 34rpx; font-weight: 600; color: #1c1c1e; }
	}
	.arrow { font-size: 36rpx; color: #c7c7cc; }
}
.menu-group {
	background: #fff; border-radius: 20rpx; margin-bottom: 24rpx; overflow: hidden;
	.menu-item {
		display: flex; justify-content: space-between; align-items: center;
		padding: 28rpx 32rpx; border-bottom: 1rpx solid #f2f2f7;
		text { font-size: 30rpx; color: #1c1c1e; }
		.arrow { font-size: 32rpx; color: #c7c7cc; }
	}
}
.logout-btn { background: #fff; color: #ff3b30; border: none; border-radius: 20rpx; font-size: 30rpx; margin-top: 20rpx; }
</style>
