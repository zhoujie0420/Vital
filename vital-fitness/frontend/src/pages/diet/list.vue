<template>
	<view class="container">
		<u-navbar title="饮食记录" fixed placeholder>
			<view slot="right">
				<text class="add-btn" @tap="addDiet">+ 记录</text>
			</view>
		</u-navbar>

		<view class="diet-list">
			<view v-for="(r, i) in dietList" :key="i" class="diet-item">
				<view class="item-header">
					<text class="item-date">{{ formatDate(r.record_date) }}</text>
					<text class="item-meal">{{ getMealName(r.meal_type) }}</text>
				</view>
				<view class="item-body">
					<text class="item-cal">{{ r.total_calories }}kcal</text>
					<text class="item-macro">蛋白{{ r.protein }}g | 碳水{{ r.carbs }}g | 脂肪{{ r.fat }}g</text>
				</view>
				<text v-if="r.notes" class="item-notes">{{ r.notes }}</text>
			</view>

			<view v-if="dietList.length === 0" class="empty-state">
				<text class="empty-text">暂无饮食记录</text>
				<button class="add-first-btn" type="primary" @tap="addDiet">记录第一餐</button>
			</view>
		</view>
	</view>
</template>

<script>
	import { getDietRecords } from '../../api/diet'

	export default {
		data() { return { dietList: [] } },
		onShow() { this.load() },
		methods: {
			async load() {
				try {
					const res = await getDietRecords({ page: 1, page_size: 50 })
					this.dietList = res.data || []
				} catch (e) {}
			},
			addDiet() { uni.navigateTo({ url: '/pages/diet/add' }) },
			getMealName(t) {
				return { breakfast: '早餐', lunch: '午餐', dinner: '晚餐', snack: '加餐' }[t] || t
			},
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
.diet-list {
	.diet-item {
		background: white; border-radius: 20rpx; padding: 30rpx; margin-bottom: 15rpx;
		.item-header { display: flex; justify-content: space-between; margin-bottom: 10rpx;
			.item-date { font-size: 24rpx; color: #999; }
			.item-meal { font-size: 24rpx; color: #3c9cff; font-weight: bold; }
		}
		.item-body {
			.item-cal { font-size: 36rpx; font-weight: bold; color: #ff9900; display: block; margin-bottom: 8rpx; }
			.item-macro { font-size: 22rpx; color: #666; }
		}
		.item-notes { font-size: 24rpx; color: #999; margin-top: 10rpx; display: block; }
	}
	.empty-state { text-align: center; padding: 100rpx 0; background: white; border-radius: 20rpx;
		.empty-text { font-size: 28rpx; color: #999; display: block; margin-bottom: 30rpx; }
		.add-first-btn { width: 60%; font-size: 28rpx; }
	}
}
</style>
