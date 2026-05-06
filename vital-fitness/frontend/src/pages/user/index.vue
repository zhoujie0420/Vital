<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="page-header">
			<text class="page-title">我的</text>
		</view>

		<!-- 已登录：用户卡片 -->
		<view v-if="isLoggedIn" class="profile-card" @tap="goTo('/pages/user/profile')">
			<view class="profile-avatar">
				<text class="avatar-text">{{ (userStore.userInfo?.nickname || '用')[0] }}</text>
			</view>
			<view class="profile-info">
				<text class="profile-name">{{ userStore.userInfo?.nickname || '微信用户' }}</text>
				<text class="profile-sub">查看和编辑个人资料</text>
			</view>
			<text class="sf-chevron">›</text>
		</view>

		<!-- 未登录：引导登录卡片 -->
		<view v-else class="profile-card guest-card" @tap="goToLogin">
			<view class="profile-avatar guest-avatar">
				<text class="avatar-text">访</text>
			</view>
			<view class="profile-info">
				<text class="profile-name">未登录</text>
				<text class="profile-sub">登录后可同步记录与数据</text>
			</view>
			<view class="login-btn-small"><text>去登录</text></view>
		</view>

		<!-- 功能菜单 -->
		<view class="menu-card">
			<view class="menu-row" @tap="goTo('/pages/user/profile')">
				<view class="menu-left">
					<view class="menu-icon bg-primary">
						<text class="menu-icon-letter">P</text>
					</view>
					<text class="menu-text">个人资料</text>
				</view>
				<text class="sf-chevron">›</text>
			</view>
			<view class="menu-row" @tap="goTo('/pages/statistics/index')">
				<view class="menu-left">
					<view class="menu-icon bg-green">
						<text class="menu-icon-letter">S</text>
					</view>
					<text class="menu-text">数据统计</text>
				</view>
				<text class="sf-chevron">›</text>
			</view>
			<view class="menu-row" @tap="goTo('/pages/diet/plan')">
				<view class="menu-left">
					<view class="menu-icon bg-teal">
						<text class="menu-icon-letter">D</text>
					</view>
					<text class="menu-text">饮食规划</text>
				</view>
				<text class="sf-chevron">›</text>
			</view>
			<view class="menu-row last" @tap="goTo('/pages/user/settings')">
				<view class="menu-left">
					<view class="menu-icon bg-gray">
						<text class="menu-icon-letter">G</text>
					</view>
					<text class="menu-text">设置</text>
				</view>
				<text class="sf-chevron">›</text>
			</view>
		</view>

		<!-- 退出登录 -->
		<view v-if="isLoggedIn" class="logout-card" @tap="logout">
			<text class="logout-text">退出登录</text>
		</view>
	</view>
</template>

<script setup>
	import { computed } from 'vue'
	import { useUserStore } from '../../store'
	import { isLoggedIn as checkLoggedIn, requireLogin } from '../../utils/authGuard'

	const userStore = useUserStore()

	const isLoggedIn = computed(() => !!userStore.token)

	const topPadding = computed(() => {
		const app = getApp()
		return (app.globalData?.customBarHeight || 88) + 8
	})

	const goToLogin = () => {
		uni.navigateTo({ url: '/pages/auth/login' })
	}

	const goTo = (url) => {
		// 未登录时，除统计页（仅本地展示）外的功能引导登录
		if (!checkLoggedIn()) {
			requireLogin()
			return
		}
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
					uni.showToast({ title: '已退出登录', icon: 'success' })
				}
			}
		})
	}
</script>

<style lang="scss" scoped>
@import '../../styles/variables.scss';

.page {
	padding: 0 $spacing-xl;
	padding-bottom: 40rpx;
	min-height: 100vh;
	background: $color-bg;
}

.page-header {
	margin-bottom: $spacing-lg;
	.page-title { @include page-title; }
}

// --- Profile Card ---
.profile-card {
	display: flex;
	align-items: center;
	@include card;
	padding: $spacing-xl;
	margin-bottom: $spacing-lg;
	@include press-effect;

	.profile-avatar {
		width: 108rpx;
		height: 108rpx;
		border-radius: 50%;
		background: $color-primary;
		display: flex;
		align-items: center;
		justify-content: center;
		margin-right: $spacing-lg;
		box-shadow: 0 4rpx 16rpx rgba(16, 185, 129, 0.25);
		flex-shrink: 0;

		.avatar-text {
			color: #fff;
			font-size: 44rpx;
			font-weight: 600;
		}
	}

	.guest-avatar {
		background: $color-separator-opaque;
		box-shadow: none;
	}

	.profile-info {
		flex: 1;

		.profile-name {
			display: block;
			font-size: $font-title3;
			font-weight: 600;
			color: $color-label;
			letter-spacing: -0.5rpx;
		}
		.profile-sub {
			display: block;
			font-size: $font-caption1;
			color: $color-label-quaternary;
			margin-top: 4rpx;
		}
	}

	.sf-chevron {
		font-size: 44rpx;
		color: $color-separator-opaque;
		font-weight: 300;
	}

	.login-btn-small {
		padding: $spacing-xs $spacing-lg;
		background: $color-primary;
		border-radius: $radius-full;
		text {
			font-size: $font-caption1;
			color: #fff;
			font-weight: 600;
		}
	}
}

// --- Menu Card ---
.menu-card {
	@include card;
	padding: 0;
	overflow: hidden;
	margin-bottom: $spacing-xl;

	.menu-row {
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
			left: 100rpx;
			right: $spacing-xl;
			height: 0.5rpx;
			background: $color-separator;
		}

		&.last::after { display: none; }

		.menu-left {
			display: flex;
			align-items: center;
			gap: $spacing-md;

			.menu-icon {
				width: 60rpx;
				height: 60rpx;
				border-radius: $radius-sm;
				display: flex;
				align-items: center;
				justify-content: center;
				flex-shrink: 0;
			}
			.menu-icon-letter {
				font-size: $font-subhead;
				font-weight: 700;
			}
			.bg-primary {
				background: $color-primary-light;
				.menu-icon-letter { color: $color-primary; }
			}
			.bg-green {
				background: $color-green-light;
				.menu-icon-letter { color: $color-green; }
			}
			.bg-teal {
				background: $color-teal-light;
				.menu-icon-letter { color: $color-teal; }
			}
			.bg-gray {
				background: $color-fill-secondary;
				.menu-icon-letter { color: $color-label-tertiary; }
			}

			.menu-text {
				font-size: $font-body;
				color: $color-label;
				font-weight: 400;
			}
		}

		.sf-chevron {
			font-size: 36rpx;
			color: $color-separator-opaque;
			font-weight: 300;
		}
	}
}

// --- Logout ---
.logout-card {
	@include card;
	text-align: center;
	padding: $spacing-lg;
	@include press-effect;

	.logout-text {
		font-size: $font-body;
		color: $color-red;
		font-weight: 500;
	}
}
</style>
