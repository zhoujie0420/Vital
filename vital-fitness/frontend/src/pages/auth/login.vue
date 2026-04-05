<template>
	<view class="container">
		<view class="header-decoration">
			<view class="decoration-circle circle-1"></view>
			<view class="decoration-circle circle-2"></view>
			<view class="decoration-circle circle-3"></view>
		</view>

		<view class="logo-section">
			<view class="logo">
				<u-icon name="level" size="80" color="#3c9cff"></u-icon>
			</view>
			<text class="app-title">Vital Fitness</text>
			<text class="app-subtitle">记录每一个进步的瞬间</text>
		</view>

		<view class="login-section">
			<button
				class="wx-login-btn"
				type="primary"
				:loading="loading"
				:disabled="loading"
				@tap="handleWxLogin"
			>
				{{ loading ? '登录中...' : '微信一键登录' }}
			</button>
		</view>
	</view>
</template>

<script>
	import { wxLogin } from '../../api/user'

	export default {
		data() {
			return {
				loading: false
			}
		},
		methods: {
			handleWxLogin() {
				this.loading = true
				uni.login({
					provider: 'weixin',
					success: async (loginRes) => {
						try {
							const res = await wxLogin(loginRes.code)
							this.$store.dispatch('saveLoginInfo', {
								token: res.data.token,
								user: res.data.user
							})
							uni.showToast({ title: '登录成功', icon: 'success' })
							setTimeout(() => {
								uni.switchTab({ url: '/pages/index/index' })
							}, 1000)
						} catch (err) {
							uni.showToast({ title: '登录失败', icon: 'none' })
						} finally {
							this.loading = false
						}
					},
					fail: () => {
						this.loading = false
						uni.showToast({ title: '微信登录取消', icon: 'none' })
					}
				})
			}
		}
	}
</script>

<style lang="scss" scoped>
.container {
	padding: 20rpx;
	min-height: 100vh;
	background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
	position: relative;
	overflow: hidden;
	display: flex;
	flex-direction: column;
	justify-content: center;
}

.header-decoration {
	.decoration-circle {
		position: absolute;
		border-radius: 50%;
		background: rgba(255, 255, 255, 0.1);
		&.circle-1 { width: 200rpx; height: 200rpx; top: -100rpx; right: -100rpx; }
		&.circle-2 { width: 300rpx; height: 300rpx; top: 200rpx; left: -150rpx; }
		&.circle-3 { width: 150rpx; height: 150rpx; bottom: 100rpx; right: 100rpx; }
	}
}

.logo-section {
	text-align: center;
	padding: 100rpx 0 80rpx;
	.logo { margin-bottom: 30rpx; }
	.app-title { font-size: 48rpx; font-weight: bold; color: white; display: block; margin-bottom: 20rpx; }
	.app-subtitle { font-size: 28rpx; color: rgba(255, 255, 255, 0.8); }
}

.login-section {
	padding: 60rpx 80rpx;

	.wx-login-btn {
		background: #3c9cff;
		color: white;
		border-radius: 20rpx;
		font-size: 32rpx;
	}
}
</style>
