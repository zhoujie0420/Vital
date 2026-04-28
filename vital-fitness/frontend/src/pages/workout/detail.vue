<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="nav-bar">
			<text class="nav-back" @tap="goBack">‹ 返回</text>
			<text class="nav-title">{{ dayLabel }}</text>
			<text class="nav-action" @tap="goAdd">+ 添加</text>
		</view>

		<view class="workout-list">
			<view v-for="(w, i) in workouts" :key="w.id" class="workout-card">
				<view class="card-top">
					<text class="card-name">{{ w.exercise.name }}</text>
					<view class="card-badge">
						<text>{{ w.exercise.category }}</text>
					</view>
				</view>
				<view class="card-stats">
					<view class="stat">
						<text class="stat-val">{{ w.weight }}</text>
						<text class="stat-unit">kg</text>
					</view>
					<view class="stat-divider"></view>
					<view class="stat">
						<text class="stat-val">{{ w.sets }}</text>
						<text class="stat-unit">组</text>
					</view>
					<view class="stat-divider"></view>
					<view class="stat">
						<text class="stat-val">{{ w.reps }}</text>
						<text class="stat-unit">次</text>
					</view>
				</view>
				<text v-if="w.notes" class="card-notes">{{ w.notes }}</text>
				<view class="card-actions">
					<text class="action-link action-edit" @tap="startEdit(w)">编辑</text>
					<text class="action-link action-delete" @tap="confirmDelete(w)">删除</text>
				</view>
			</view>
		</view>

		<view v-if="workouts.length === 0" class="empty">
			<view class="empty-icon">
				<text class="empty-icon-letter">W</text>
			</view>
			<text class="empty-title">当天暂无训练记录</text>
			<text class="empty-desc">点击右上角添加训练</text>
		</view>

		<!-- 编辑弹窗 -->
		<view class="mask" v-if="showEdit" @tap="showEdit = false"></view>
		<view class="edit-sheet" v-if="showEdit">
			<view class="sheet-handle"></view>
			<text class="sheet-title">编辑训练</text>

			<view class="edit-row-group">
				<view class="edit-row">
					<text class="edit-label">重量</text>
					<view class="edit-input-wrap">
						<input type="digit" v-model="editForm.weight" class="edit-input" />
						<text class="edit-suffix">kg</text>
					</view>
				</view>
				<view class="edit-row">
					<text class="edit-label">组数</text>
					<view class="edit-input-wrap">
						<input type="number" v-model="editForm.sets" class="edit-input" />
						<text class="edit-suffix">组</text>
					</view>
				</view>
				<view class="edit-row">
					<text class="edit-label">次数</text>
					<view class="edit-input-wrap">
						<input type="number" v-model="editForm.reps" class="edit-input" />
						<text class="edit-suffix">次</text>
					</view>
				</view>
			</view>

			<view class="sheet-field">
				<text class="sheet-label">备注</text>
				<view class="sheet-input-wrap">
					<input v-model="editForm.notes" placeholder="可选" class="sheet-input" />
				</view>
			</view>

			<view class="sheet-actions">
				<view class="sheet-cancel" @tap="showEdit = false">
					<text>取消</text>
				</view>
				<view class="sheet-save" @tap="saveEdit">
					<text>保存</text>
				</view>
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
		onLoad(opts) { this.date = opts.date || '' },
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
@import '../../styles/variables.scss';

.page {
	padding: 0 $spacing-xl;
	padding-bottom: 40rpx;
	min-height: 100vh;
	background: $color-bg;
}

.nav-bar {
	@include nav-bar;
	.nav-back { font-size: $font-callout; color: $color-primary; font-weight: 500; }
	.nav-title { font-size: $font-headline; font-weight: 600; color: $color-label; }
	.nav-action { font-size: $font-subhead; color: $color-primary; font-weight: 600; }
}

// --- Workout Card ---
.workout-card {
	@include card;
	margin-bottom: $spacing-sm;

	.card-top {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: $spacing-md;

		.card-name {
			font-size: $font-headline;
			font-weight: 600;
			color: $color-label;
		}
		.card-badge {
			padding: 6rpx $spacing-md;
			border-radius: $spacing-xs;
			font-size: $font-caption1;
			font-weight: 500;
			background: $color-primary-light;
			color: $color-primary;
		}
	}

	.card-stats {
		display: flex;
		align-items: center;
		gap: $spacing-lg;
		margin-bottom: $spacing-md;

		.stat {
			display: flex;
			align-items: baseline;
			gap: 4rpx;

			.stat-val {
				font-size: 44rpx;
				font-weight: 700;
				color: $color-label;
				font-variant-numeric: tabular-nums;
			}
			.stat-unit {
				font-size: $font-caption1;
				color: $color-label-quaternary;
			}
		}

		.stat-divider {
			width: 1rpx;
			height: 32rpx;
			background: $color-separator;
		}
	}

	.card-notes {
		display: block;
		font-size: $font-footnote;
		color: $color-label-quaternary;
		margin-bottom: $spacing-md;
		line-height: 1.5;
	}

	.card-actions {
		display: flex;
		gap: $spacing-lg;
		padding-top: $spacing-md;
		border-top: 0.5rpx solid $color-separator;

		.action-link {
			font-size: $font-subhead;
			font-weight: 500;
			padding: $spacing-xs 0;

			&.action-edit { color: $color-primary; }
			&.action-delete { color: $color-red; }
			&:active { opacity: 0.5; }
		}
	}
}

.empty {
	@include empty-state;

	.empty-icon-letter {
		font-size: 44rpx;
		font-weight: 700;
		color: $color-label-quaternary;
	}
}

// --- Edit Sheet ---
.mask {
	position: fixed;
	top: 0; left: 0; right: 0; bottom: 0;
	background: rgba(0, 0, 0, 0.35);
	z-index: 100;
}

.edit-sheet {
	position: fixed;
	bottom: 0;
	left: 0;
	right: 0;
	background: $color-bg-elevated;
	border-radius: $radius-xl $radius-xl 0 0;
	padding: $spacing-md $spacing-xl $spacing-2xl;
	padding-bottom: calc(#{$spacing-2xl} + env(safe-area-inset-bottom));
	z-index: 101;

	.sheet-handle {
		width: 72rpx;
		height: 8rpx;
		background: $color-fill-tertiary;
		border-radius: 4rpx;
		margin: 0 auto $spacing-lg;
	}

	.sheet-title {
		display: block;
		font-size: $font-title3;
		font-weight: 700;
		color: $color-label;
		text-align: center;
		margin-bottom: $spacing-xl;
	}
}

// --- Edit Row Group ---
.edit-row-group {
	display: flex;
	gap: $spacing-sm;
	margin-bottom: $spacing-lg;
}

.edit-row {
	flex: 1;

	.edit-label {
		display: block;
		font-size: $font-caption2;
		color: $color-label-quaternary;
		font-weight: 500;
		margin-bottom: $spacing-xs;
		text-align: center;
	}
}

.edit-input-wrap {
	display: flex;
	align-items: center;
	justify-content: center;
	background: $color-fill;
	border-radius: $radius-md;
	padding: $spacing-md $spacing-sm;
	border: 1.5rpx solid transparent;
	transition: all 0.2s ease;

	&:focus-within {
		background: rgba(16, 185, 129, 0.04);
		border-color: rgba(16, 185, 129, 0.3);
	}

	.edit-input {
		width: 80rpx;
		text-align: center;
		font-size: $font-headline;
		font-weight: 700;
		color: $color-label;
		font-variant-numeric: tabular-nums;
	}
	.edit-suffix {
		font-size: $font-caption1;
		color: $color-label-quaternary;
		font-weight: 500;
		margin-left: 4rpx;
	}
}

// --- Notes & Actions ---
.sheet-field {
	margin-bottom: $spacing-md;

	.sheet-label {
		display: block;
		font-size: $font-footnote;
		color: $color-label-quaternary;
		font-weight: 500;
		margin-bottom: $spacing-xs;
	}
	.sheet-input-wrap {
		@include input-field;
	}
	.sheet-input {
		flex: 1;
		font-size: $font-body;
		color: $color-label;
	}
}

.sheet-actions {
	display: flex;
	gap: $spacing-md;
	margin-top: $spacing-md;

	.sheet-cancel, .sheet-save {
		flex: 1;
		text-align: center;
		padding: $spacing-lg 0;
		border-radius: $radius-md;
		font-size: $font-callout;
		font-weight: 600;
	}
	.sheet-cancel {
		background: $color-fill;
		color: $color-label-tertiary;
	}
	.sheet-save {
		@include primary-button;
	}
}
</style>
