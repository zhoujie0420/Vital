<template>
	<view class="page">
		<view class="header">
			<view class="greeting">
				<text class="greeting-sub">{{ greetingTime }}</text>
				<text class="greeting-name">{{ userStore.userInfo?.nickname || '健身达人' }}</text>
			</view>
		</view>

		<view class="summary-card">
			<view class="summary-row">
				<view class="summary-item">
					<text class="summary-value">{{ dashboard.today_workouts || 0 }}</text>
					<text class="summary-label">训练</text>
				</view>
				<view class="summary-divider"></view>
				<view class="summary-item">
					<text class="summary-value">{{ dashboard.today_calories || 0 }}</text>
					<text class="summary-label">千卡</text>
				</view>
				<view class="summary-divider"></view>
				<view class="summary-item">
					<text class="summary-value">{{ dashboard.latest_weight || '--' }}</text>
					<text class="summary-label">体重kg</text>
				</view>
			</view>
		</view>

		<view class="section">
			<text class="section-title">快捷记录</text>
			<view class="action-grid">
				<view class="action-card" @tap="goTo('/pages/workout/add')"><text class="action-text">训练</text></view>
				<view class="action-card" @tap="goTo('/pages/diet/add')"><text class="action-text">饮食</text></view>
				<view class="action-card" @tap="goTo('/pages/weight/add')"><text class="action-text">体重</text></view>
				<view class="action-card" @tap="goTo('/pages/mood/add')"><text class="action-text">心情</text></view>
			</view>
		</view>

		<view class="section">
			<text class="section-title">数据概览</text>
			<view class="data-card" @tap="goTo('/pages/workout/list')">
				<text class="data-title">训练次数</text>
				<text class="data-value">{{ dashboard.workout_count || 0 }}</text>
			</view>
		</view>
	</view>
</template>

<script setup>
	import { ref, computed, onMounted } from 'vue'
	import { useUserStore } from '../../store'
	import { getDashboard } from '../../api/stats'

	const userStore = useUserStore()
	const dashboard = ref({})

	const greetingTime = computed(() => {
		const h = new Date().getHours()
		if (h < 6) return '夜深了'
		if (h < 12) return '早上好'
		if (h < 18) return '下午好'
		return '晚上好'
	})

	const goTo = (url) => uni.navigateTo({ url })

	const loadDashboard = async () => {
		try {
			const res = await getDashboard()
			dashboard.value = res.data || {}
		} catch (e) {}
	}

	onMounted(() => loadDashboard())
</script>

<style lang="scss" scoped>
.page { padding: 32rpx; padding-top: 120rpx; min-height: 100vh; background: #f2f2f7; }
.header { margin-bottom: 40rpx;
	.greeting-sub { display: block; font-size: 28rpx; color: #8e8e93; }
	.greeting-name { display: block; font-size: 48rpx; font-weight: 700; color: #1c1c1e; }
}
.summary-card {
	background: #fff; border-radius: 20rpx; padding: 36rpx 0; margin-bottom: 40rpx;
	.summary-row { display: flex; align-items: center; }
	.summary-item { flex: 1; text-align: center;
		.summary-value { display: block; font-size: 44rpx; font-weight: 700; color: #1c1c1e; }
		.summary-label { display: block; font-size: 24rpx; color: #8e8e93; margin-top: 6rpx; }
	}
	.summary-divider { width: 1rpx; height: 60rpx; background: #e5e5ea; }
}
.section { margin-bottom: 40rpx;
	.section-title { display: block; font-size: 36rpx; font-weight: 700; color: #1c1c1e; margin-bottom: 20rpx; }
}
.action-grid { display: flex; gap: 16rpx;
	.action-card { flex: 1; background: #fff; border-radius: 20rpx; padding: 28rpx 0; text-align: center;
		.action-text { font-size: 28rpx; color: #3a3a3c; font-weight: 500; }
	}
}
.data-card {
	display: flex; justify-content: space-between; align-items: center;
	background: #fff; border-radius: 16rpx; padding: 32rpx;
	.data-title { font-size: 30rpx; font-weight: 600; color: #1c1c1e; }
	.data-value { font-size: 40rpx; font-weight: 700; color: #007aff; }
}
</style>
