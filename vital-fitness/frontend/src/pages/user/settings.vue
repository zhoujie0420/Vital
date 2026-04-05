<template>
	<view class="container">
		<view class="nav-bar">
			<text class="back" @tap="goBack">‹ 返回</text>
			<text class="nav-title">设置</text>
		</view>

		<view class="menu-group">
			<view class="menu-item">
				<text>版本</text><text class="value">v1.0.0</text>
			</view>
			<view class="menu-item" @tap="clearCache">
				<text>清除缓存</text><text class="arrow">›</text>
			</view>
		</view>
	</view>
</template>

<script setup>
	const goBack = () => uni.navigateBack()
	const clearCache = () => {
		uni.showModal({
			title: '清除缓存', content: '确定清除？',
			success: (res) => {
				if (res.confirm) uni.showToast({ title: '已清除', icon: 'success' })
			}
		})
	}
</script>

<style lang="scss" scoped>
.container { padding: 32rpx; }
.nav-bar { display: flex; align-items: center; gap: 20rpx; margin-bottom: 20rpx;
	.back { font-size: 32rpx; color: #007aff; }
	.nav-title { font-size: 36rpx; font-weight: 700; color: #1c1c1e; }
}
.menu-group {
	background: #fff; border-radius: 20rpx; overflow: hidden;
	.menu-item {
		display: flex; justify-content: space-between; align-items: center;
		padding: 28rpx 32rpx; border-bottom: 1rpx solid #f2f2f7;
		text { font-size: 30rpx; color: #1c1c1e; }
		.value { color: #8e8e93; }
		.arrow { font-size: 32rpx; color: #c7c7cc; }
	}
}
</style>
