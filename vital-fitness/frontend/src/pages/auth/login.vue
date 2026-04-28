<template>
	<view class="page">
		<view class="bg-layer">
			<view class="bg-orb orb-1"></view>
			<view class="bg-orb orb-2"></view>
		</view>

		<view class="content">
			<view class="logo-section">
				<view class="logo-mark">
					<text class="logo-letter">V</text>
				</view>
				<text class="app-title">Vital Fitness</text>
				<text class="app-subtitle">记录每一个进步的瞬间</text>
			</view>

			<view class="login-section">
				<view class="login-btn" :class="{ loading: loading }" @tap="handleWxLogin">
					<text class="btn-text">{{ loading ? '登录中...' : '微信一键登录' }}</text>
				</view>
				<text class="login-hint">使用微信账号安全登录</text>
			</view>
		</view>
	</view>
</template>

<script setup>
	import { ref } from 'vue'
	import { useUserStore } from '../../store'
	import { wxLogin } from '../../api/user'

	const loading = ref(false)
	const userStore = useUserStore()

	const handleWxLogin = () => {
		loading.value = true
		uni.login({
			provider: 'weixin',
			success: async (loginRes) => {
				try {
					const res = await wxLogin(loginRes.code)
					userStore.saveLoginInfo({ token: res.data.token, user: res.data.user })
					uni.showToast({ title: '登录成功', icon: 'success' })
					setTimeout(() => uni.switchTab({ url: '/pages/index/index' }), 1000)
				} catch (err) {
					uni.showToast({ title: '登录失败', icon: 'none' })
				} finally { loading.value = false }
			},
			fail: () => {
				loading.value = false
				uni.showToast({ title: '微信登录取消', icon: 'none' })
			}
		})
	}
</script>

<style lang="scss" scoped>
@import '../../styles/variables.scss';

.page {
	min-height: 100vh;
	background: #111110;
	position: relative;
	overflow: hidden;
}

.bg-layer {
	position: absolute;
	top: 0; left: 0; right: 0; bottom: 0;

	.bg-orb {
		position: absolute;
		border-radius: 50%;
		filter: blur(120rpx);
	}
	.orb-1 {
		width: 520rpx; height: 520rpx;
		background: rgba(16, 185, 129, 0.20);
		top: -100rpx; right: -120rpx;
	}
	.orb-2 {
		width: 480rpx; height: 480rpx;
		background: rgba(20, 184, 166, 0.12);
		bottom: 80rpx; left: -160rpx;
	}
}

.content {
	position: relative;
	z-index: 1;
	min-height: 100vh;
	display: flex;
	flex-direction: column;
	justify-content: center;
	padding: 60rpx 48rpx;
}

.logo-section {
	text-align: center;
	margin-bottom: 120rpx;

	.logo-mark {
		width: 140rpx;
		height: 140rpx;
		border-radius: $radius-2xl;
		background: $color-primary;
		display: flex;
		align-items: center;
		justify-content: center;
		margin: 0 auto 36rpx;
		box-shadow: 0 8rpx 40rpx rgba(16, 185, 129, 0.3);

		.logo-letter {
			font-size: 64rpx;
			font-weight: 700;
			color: #fff;
			letter-spacing: -2rpx;
		}
	}

	.app-title {
		display: block;
		font-size: 52rpx;
		font-weight: 700;
		color: #FAFAF9;
		letter-spacing: -1.5rpx;
		margin-bottom: $spacing-sm;
	}
	.app-subtitle {
		display: block;
		font-size: $font-subhead;
		color: rgba(250, 250, 249, 0.35);
		font-weight: 400;
		letter-spacing: 0.5rpx;
	}
}

.login-section {
	.login-btn {
		background: rgba(250, 250, 249, 0.08);
		border: 0.5rpx solid rgba(250, 250, 249, 0.10);
		backdrop-filter: saturate(180%) blur(40px);
		-webkit-backdrop-filter: saturate(180%) blur(40px);
		border-radius: $radius-lg;
		padding: $spacing-xl 0;
		text-align: center;
		transition: all 0.25s cubic-bezier(0.16, 1, 0.3, 1);

		&:active {
			transform: scale(0.97);
			background: rgba(250, 250, 249, 0.14);
		}
		&.loading { opacity: 0.5; pointer-events: none; }

		.btn-text {
			font-size: $font-callout;
			font-weight: 600;
			color: #FAFAF9;
		}
	}

	.login-hint {
		display: block;
		text-align: center;
		font-size: $font-caption1;
		color: rgba(250, 250, 249, 0.2);
		margin-top: $spacing-lg;
	}
}
</style>
