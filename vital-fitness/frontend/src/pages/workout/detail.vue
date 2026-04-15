<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="nav-bar">
			<text class="nav-back" @tap="goBack">‹ 返回</text>
			<text class="nav-title">{{ dayLabel }}</text>
			<text class="nav-action" @tap="goAdd">+ 添加</text>
		</view>

		<view class="workout-list">
			<view v-for="(w, i) in workouts" :key="w.id" class="workout-card">
				<view class="card-header">
					<text class="card-name">{{ w.exercise.name }}</text>
					<text class="card-category">{{ w.exercise.category }}</text>
				</view>
				<view class="card-stats">
					<view class="stat">
						<text class="stat-val">{{ w.weight }}</text>
						<text class="stat-unit">kg</text>
					</view>
					<view class="stat">
						<text class="stat-val">{{ w.sets }}</text>
						<text class="stat-unit">组</text>
					</view>
					<view class="stat">
						<text class="stat-val">{{ w.reps }}</text>
						<text class="stat-unit">次</text>
					</view>
				</view>
				<text v-if="w.notes" class="card-notes">{{ w.notes }}</text>
				<view class="card-actions">
					<text class="action-btn edit" @tap="startEdit(w)">编辑</text>
					<text class="action-btn delete" @tap="confirmDelete(w)">删除</text>
				</view>
			</view>
		</view>

		<view v-if="workouts.length === 0" class="empty">
			<text class="empty-text">当天暂无训练记录</text>
		</view>

		<!-- 编辑弹窗 -->
		<view class="mask" v-if="showEdit" @tap="showEdit = false"></view>
		<view class="edit-popup" v-if="showEdit">
			<text class="popup-title">编辑训练</text>
			<view class="popup-field">
				<text class="popup-label">重量 (kg)</text>
				<input type="digit" v-model="editForm.weight" class="popup-input" />
			</view>
			<view class="popup-field">
				<text class="popup-label">组数</text>
				<input type="number" v-model="editForm.sets" class="popup-input" />
			</view>
			<view class="popup-field">
				<text class="popup-label">次数</text>
				<input type="number" v-model="editForm.reps" class="popup-input" />
			</view>
			<view class="popup-field">
				<text class="popup-label">备注</text>
				<input v-model="editForm.notes" placeholder="可选" class="popup-input" />
			</view>
			<view class="popup-actions">
				<text class="popup-cancel" @tap="showEdit = false">取消</text>
				<text class="popup-save" @tap="saveEdit">保存</text>
			</view>
		</view>
	</view>
</template>

<script>
	import { getWorkouts, updateWorkout, deleteWorkout } from '../../api/workout'

	export default {
		data() {
			return {
				date: '',
				workouts: [],
				showEdit: false,
				editingId: null,
				editForm: { weight: '', sets: '', reps: '', notes: '' }
			}
		},
		computed: {
			topPadding() {
				const app = getApp()
				return (app.globalData?.customBarHeight || 88) + 8
			},
			dayLabel() {
				if (!this.date) return ''
				const d = new Date(this.date + 'T00:00:00')
				const today = new Date()
				const yesterday = new Date(today)
				yesterday.setDate(yesterday.getDate() - 1)
				if (d.toDateString() === today.toDateString()) return '今天的训练'
				if (d.toDateString() === yesterday.toDateString()) return '昨天的训练'
				return (d.getMonth() + 1) + '月' + d.getDate() + '日训练'
			}
		},
		onLoad(opts) {
			this.date = opts.date || ''
		},
		onShow() { this.load() },
		methods: {
			goBack() { uni.navigateBack() },
			goAdd() { uni.navigateTo({ url: '/pages/workout/add' }) },
			async load() {
				if (!this.date) return
				try {
					const res = await getWorkouts({ page: 1, page_size: 200 })
					const all = res.data || []
					this.workouts = all.filter(w => {
						const d = new Date(w.workout_date)
						return d.toISOString().split('T')[0] === this.date
					})
				} catch (e) {}
			},
			startEdit(w) {
				this.editingId = w.id
				this.editForm = {
					weight: String(w.weight),
					sets: String(w.sets),
					reps: String(w.reps),
					notes: w.notes || ''
				}
				this.showEdit = true
			},
			async saveEdit() {
				try {
					await updateWorkout(this.editingId, {
						weight: parseFloat(this.editForm.weight),
						sets: parseInt(this.editForm.sets),
						reps: parseInt(this.editForm.reps),
						notes: this.editForm.notes
					})
					uni.showToast({ title: '更新成功', icon: 'success' })
					this.showEdit = false
					this.load()
				} catch (e) {}
			},
			confirmDelete(w) {
				uni.showModal({
					title: '确认删除',
					content: `删除「${w.exercise.name}」这条记录？`,
					success: async (res) => {
						if (res.confirm) {
							try {
								await deleteWorkout(w.id)
								uni.showToast({ title: '已删除', icon: 'success' })
								this.load()
							} catch (e) {}
						}
					}
				})
			}
		}
	}
</script>

<style lang="scss" scoped>
.page {
	padding: 0 32rpx;
	padding-bottom: 40rpx;
	min-height: 100vh;
	background: #f2f2f7;
}

.nav-bar {
	display: flex;
	align-items: center;
	justify-content: space-between;
	margin-bottom: 28rpx;

	.nav-back { font-size: 32rpx; color: #007aff; font-weight: 500; }
	.nav-title { font-size: 34rpx; font-weight: 600; color: #1c1c1e; }
	.nav-action { font-size: 30rpx; color: #007aff; font-weight: 600; }
}

.workout-card {
	background: #fff;
	border-radius: 20rpx;
	padding: 28rpx 32rpx;
	margin-bottom: 16rpx;
	box-shadow: 0 2rpx 16rpx rgba(0, 0, 0, 0.04);

	.card-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 20rpx;

		.card-name {
			font-size: 34rpx;
			font-weight: 600;
			color: #1c1c1e;
		}
		.card-category {
			font-size: 24rpx;
			color: #007aff;
			background: rgba(0, 122, 255, 0.08);
			padding: 6rpx 16rpx;
			border-radius: 10rpx;
			font-weight: 500;
		}
	}

	.card-stats {
		display: flex;
		gap: 40rpx;
		margin-bottom: 16rpx;

		.stat {
			display: flex;
			align-items: baseline;
			gap: 4rpx;

			.stat-val {
				font-size: 40rpx;
				font-weight: 700;
				color: #1c1c1e;
			}
			.stat-unit {
				font-size: 24rpx;
				color: #8e8e93;
			}
		}
	}

	.card-notes {
		display: block;
		font-size: 26rpx;
		color: #8e8e93;
		margin-bottom: 16rpx;
		line-height: 1.5;
	}

	.card-actions {
		display: flex;
		gap: 24rpx;
		padding-top: 16rpx;
		border-top: 1rpx solid #f2f2f7;

		.action-btn {
			font-size: 28rpx;
			font-weight: 500;
			padding: 8rpx 0;

			&.edit { color: #007aff; }
			&.delete { color: #ff3b30; }
			&:active { opacity: 0.6; }
		}
	}
}

.empty {
	text-align: center;
	padding: 80rpx 0;

	.empty-text { font-size: 28rpx; color: #8e8e93; }
}

.mask {
	position: fixed;
	top: 0; left: 0; right: 0; bottom: 0;
	background: rgba(0, 0, 0, 0.4);
	z-index: 100;
}

.edit-popup {
	position: fixed;
	bottom: 0;
	left: 0;
	right: 0;
	background: #fff;
	border-radius: 24rpx 24rpx 0 0;
	padding: 40rpx 32rpx;
	padding-bottom: calc(40rpx + env(safe-area-inset-bottom));
	z-index: 101;

	.popup-title {
		display: block;
		font-size: 36rpx;
		font-weight: 700;
		color: #1c1c1e;
		text-align: center;
		margin-bottom: 32rpx;
	}

	.popup-field {
		margin-bottom: 20rpx;

		.popup-label {
			display: block;
			font-size: 26rpx;
			color: #8e8e93;
			font-weight: 500;
			margin-bottom: 8rpx;
		}
		.popup-input {
			background: #f2f2f7;
			border-radius: 14rpx;
			padding: 22rpx;
			font-size: 30rpx;
			color: #1c1c1e;
		}
	}

	.popup-actions {
		display: flex;
		gap: 20rpx;
		margin-top: 24rpx;

		.popup-cancel, .popup-save {
			flex: 1;
			text-align: center;
			padding: 28rpx 0;
			border-radius: 16rpx;
			font-size: 32rpx;
			font-weight: 600;
		}
		.popup-cancel {
			background: #f2f2f7;
			color: #636366;
		}
		.popup-save {
			background: #007aff;
			color: #fff;
			box-shadow: 0 4rpx 16rpx rgba(0, 122, 255, 0.3);

			&:active { transform: scale(0.98); opacity: 0.9; }
		}
	}
}
</style>
