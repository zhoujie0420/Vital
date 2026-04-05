<template>
	<view class="container">
		<u-navbar title="体重记录" fixed placeholder>
			<view slot="right">
				<text class="add-btn" @tap="addWeight">+ 记录</text>
			</view>
		</u-navbar>

		<view class="weight-list">
			<view v-for="(r, i) in weightList" :key="i" class="weight-item">
				<view class="item-header">
					<text class="item-date">{{ formatDate(r.record_date) }}</text>
				</view>
				<view class="item-body">
					<text class="item-weight">{{ r.weight }}kg</text>
					<text v-if="r.bmi" class="item-bmi">BMI: {{ r.bmi }}</text>
				</view>
			</view>

			<view v-if="weightList.length === 0" class="empty-state">
				<text class="empty-text">暂无体重记录</text>
				<button class="add-first-btn" type="primary" @tap="addWeight">记录第一次体重</button>
			</view>
		</view>
	</view>
</template>

<script>
	import { getWeights } from '../../api/weight'

	export default {
		data() { return { weightList: [] } },
		onShow() { this.load() },
		methods: {
			async load() {
				try {
					const res = await getWeights({ page: 1, page_size: 50 })
					this.weightList = res.data || []
				} catch (e) {}
			},
			addWeight() { uni.navigateTo({ url: '/pages/weight/add' }) },
			formatDate(d) {
				if (!d) return ''
				const date = new Date(d)
				return (date.getMonth()+1) + '月' + date.getDate() + '日'
			}
		}
	}
</script>

<style lang="scss" scoped>
.container { padding: 20rpx; }
.add-btn { color: #3c9cff; font-size: 28rpx; font-weight: bold; }
.weight-list {
	.weight-item {
		background: white; border-radius: 20rpx; padding: 30rpx; margin-bottom: 15rpx;
		.item-header { margin-bottom: 10rpx;
			.item-date { font-size: 24rpx; color: #999; }
		}
		.item-body { display: flex; align-items: baseline; gap: 20rpx;
			.item-weight { font-size: 40rpx; font-weight: bold; color: #333; }
			.item-bmi { font-size: 24rpx; color: #3c9cff; }
		}
	}
	.empty-state { text-align: center; padding: 100rpx 0; background: white; border-radius: 20rpx;
		.empty-text { font-size: 28rpx; color: #999; display: block; margin-bottom: 30rpx; }
		.add-first-btn { width: 60%; font-size: 28rpx; }
	}
}
</style>
