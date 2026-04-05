<template>
	<view class="container">
		<text>请使用微信登录</text>
		<u-button type="primary" @click="toLogin">去登录</u-button>
	</view>
</template>

<script>
	export default {
		methods: {
			toLogin() {
				uni.redirectTo({ url: '/pages/auth/login' })
			}
		}
	}
</script>

<style lang="scss" scoped>
.container {
	padding: 40rpx;
	text-align: center;
}
</style>
