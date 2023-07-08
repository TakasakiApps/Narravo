<template>
	<div style="margin: 0; padding: 0;">
		<el-carousel height="93vh" direction="horizontal" :autoplay="false" arrow="never" indicator-position="none"
			ref="carousel">
			<el-carousel-item :key="1" class="mian">
				<div v-if="true">
					<h1 class="mian-text">
						暂无角色
						<div>点击右下角按钮以创建
						</div>
					</h1>
				</div>
				<sel-affix position="bottom" :offset="50" class="fixed-element" @click="addRoleCard" v-show="button[0]">
					<svg height="48" viewBox="-250 -1200 1460 1460" width="48" class="btns radiusBtn">
						<g>
							<path
								d="M479.825-190.391q-16.782 0-28.108-11.451T440.391-230v-210.391H230q-16.707 0-28.158-11.501-11.451-11.502-11.451-28.283 0-16.782 11.451-28.108T230-519.609h210.391V-730q0-16.707 11.501-28.158 11.502-11.451 28.283-11.451 16.782 0 28.108 11.451T519.609-730v210.391H730q16.707 0 28.158 11.501 11.451 11.502 11.451 28.283 0 16.782-11.451 28.108T730-440.391H519.609V-230q0 16.707-11.501 28.158-11.502 11.451-28.283 11.451Z"
								fill="currentColor" />
						</g>
					</svg>
				</sel-affix>
			</el-carousel-item>
			<el-carousel-item :key="2">
				<div>
					<el-card class="box-card">
						<div class="guide">
							<transition name="guideFst">
								<div key="existingElement" v-show="!addTxtIndex">
									<div class="guide-child">
										创建您的角色
									</div>
									<div class="guide-child">
										作为小说朗读者
									</div>
								</div>
							</transition>
							<transition name="guideSed">
								<div v-show="addTxtIndex" key="newElement">
									<div class="guide-child">
										好的。现在来为
										<div>您的角色调整参数</div>
									</div>
								</div>
							</transition>
						</div>
						<el-card class="box-card-main">
							<el-carousel height="65vh" direction="vertical" :autoplay="false" arrow="never" ref="card"
								indicator-position="none">
								<el-carousel-item :key="2 - 1">
									<el-avatar :size="160" class="avatar" />
									<div class="Roleinput">
										<el-input v-model="Card.name" placeholder="请输入你想要创建的角色名" />
										<div class="avatar-txt">
											<div>通过创建自定义角色来个性化有声小说</div>
											<div>跟随向导以创建您的自定义角色</div>
										</div>
									</div>
								</el-carousel-item>
								<el-carousel-item :key="2 - 2">
									<div class="carValue">
										<div class="cardValue-title">语言模型</div>
										<el-select v-model="Card.Reader" class="cardInput" placeholder="Select" size="default">
											<el-option label="成男1号姬" value="成男1" />
										</el-select>
									</div>
									<div>
										<div class="cardValue-title">默认情感</div>
										<el-select v-model="Card" class="cardInput" placeholder="Select" size="default">
											<el-option label="不支持" value="no" />
										</el-select>
									</div>
									<div>
										<div class="cardValue-title">情感等级</div>
										<el-slider v-model="Card.emotion" class="cardInput" :step="1" size="small" show-stops :max="200"
											:disabled="canEmotion" />
									</div>
									<div>
										<div class="cardValue-title">语速</div>
										<el-slider v-model="Card.speed" :step="0.5" class="cardInput" size="small" show-stops :max="2.5"
											:min="-2" />
									</div>
									<div>
										<div class="cardValue-title">音量</div>
										<el-slider v-model="Card.tone" :step="1" class="cardInput" size="small" show-stops :max="10" />
									</div>
								</el-carousel-item>
							</el-carousel>
						</el-card>
					</el-card>
					<div @click="backBtn">
						<svg height="48" viewBox="-30 -750 600 900" width="48" class="btns radiusBtn backBtn">
							<g transform="scale(0.6)">
								<path d="M480-160 160-480l320-320 42 42-248 248h526v60H274l248 248-42 42Z" fill="currentColor" />
							</g>
						</svg>
					</div>
					<div>
						<div v-show="!addTxtIndex" @click="alterChild">
							<svg height="48" viewBox="-30 -750 600 900" width="48" class="btns radiusBtn fsBtn">
								<g transform="scale(0.6)">
									<path
										d="M452.217-183.13q-11.391-10.827-11.391-27.501t11.391-28.065l201.694-201.695h-460.52q-16.956 0-28.282-11.326-11.327-11.326-11.327-28.283t11.327-28.283q11.326-11.326 28.282-11.326h460.52l-201.694-202.26q-11.391-10.826-11.391-28t11.391-28.566q11.392-10.826 28.066-10.826t28.065 10.826L779-507.783q6.13 6.131 8.978 13.109T790.826-480q0 7.261-2.848 14.457-2.848 7.195-8.978 13.326L508.348-181.565q-11.391 11.391-28.065 10.609-16.674-.783-28.066-12.174Z"
										fill="currentColor" />
								</g>
							</svg>
						</div>
						<div v-show="addTxtIndex">
							<svg height="48" viewBox="-30 -750 600 900" width="48" class="btns radiusBtn fsBtn">
								<g transform="scale(0.6)">
									<path v-show="!addTxtIndex" d="M378-246 154-470l43-43 181 181 384-384 43 43-427 427Z"
										fill="currentColor" />
								</g>
							</svg>
						</div>

					</div>
				</div>
			</el-carousel-item>

		</el-carousel>
	</div>
</template>
    
<script setup lang='ts'>
// import RoleCard from '../components/roleCards/RoleCard.vue';
import readRoledir from '../components/roleCards/node-api'
import { onMounted, reactive, ref } from 'vue';
import { StarFilled, Plus } from '@element-plus/icons-vue'

onMounted(() => {
	// const e = ReadDir()
	// if (Array.isArray(e)) {
	// 	// 对数组类型进行断言
	// 	// 使用 s.length 进行进一步的验证

	// 	if (e.length === 0) {
	// 		// console.log('没有角色卡');
	// 	} else {
	// 		// console.log('具有角色卡');
	// 		console.log(e)
	// 	}
	// } else {
	// 	// 对其他类型进行处理（例如字符串或布尔值）
	// 	console.log('出现意料之外的错误')
	// }
})
const Card = reactive({
	name: '',
	rgb: '111',
	Reader: '小帅',
	speed: 0,
	tone: 5,
	emotion: 100
})
const button = reactive([
	true,
	false,
])
const canEmotion = ref('')
const carousel = ref<HTMLFormElement>()
function buttonShow() {
	button[0] = !button[0];

}

function addRoleCard() {
	if (carousel.value) {
		setTimeout(buttonShow, 500)
		carousel.value.next()
	}

}
function backBtn() {
	if (carousel.value) {
		buttonShow()
		carousel.value.prev()
	}
}
function alterChild() {
	addTxtIndex.value = true
	console.log(addTxtIndex.value)
}
const addTxtIndex = ref(false)
</script>
    
<style scoped>
.radiusBtn {
	box-shadow: 1px 1px 3px rgba(71, 56, 56, 0.661);
	color: white;
	border-radius: 100px;
}

.fixed-element {
	/* position: absolute;
	bottom: 20px;
	right: 20px; */
	/* display: flex;
	float: left; */
	margin-top: auto;
	margin-left: auto;
	margin-right: 5px;
	/* transform: translateX(50%); */
}

.fixed-element :hover {
	/* transform: rotate(180deg); */
}

.backBtn {
	position: absolute;
	left: 1%;
	bottom: 5%;
}

.fsBtn {
	position: absolute;
	right: 3%;
	bottom: 5%;
}

.btns {}



.mian-text {
	position: absolute;
	right: 35vw;
	top: 45vh;
	text-align: center;
	letter-spacing: 1px;
	color: rgb(192, 190, 190);
}

.mian {
	display: flex;
	flex-direction: column;
	justify-content: flex-end;
	align-items: flex-end;
}

/*  */
.box-card {
	position: absolute;
	width: 70vw;
	height: 71vh;
	left: 12%;
	margin-top: 50px;
	background-image: linear-gradient(to right top, #3a586b, #405d7a, #4d6288, #606593, #77669b, #93699e, #ad6c9c, #c57097, #dd7d8e, #ed8f85, #f6a47e, #f6bc7d);
}

.box-card-main {
	width: calc(40vw/1.618);
	height: 65vh;
	/* margin-left: 50vh; */
	position: absolute;
	/* margin: auto; */
	right: 20px;

}

/* 引导文字 */
.guide {
	position: absolute;
	left: 15%;
	top: 40%;
	font-size: 3.2vw;
	font-weight: 900;
	text-align: center;
	/* background-image: -webkit-linear-gradient(right top, #cdd1d3, #c9d3d8, #c5d6dc, #c0d8df, #bbdbe2, #bbdbe2, #bbdbe2, #bbdbe2, #c0d8df, #c5d6dc, #c9d3d8, #cdd1d3);
	-webkit-background-clip: text; */
	/* -webkit-text-fill-color: transparent; */
	color: #cfddf5;
}

.guide-child {
	display: block;
}

.avatar {
	position: relative;
	/* left: 15%; */
	margin-left: calc(50% - 80px);
	/* top: 50px; */
}

.Roleinput {
	position: relative;
	top: calc(65vh /3.5);
	width: 18vw;
	margin-left: 5%;
}

.avatar-txt {
	text-align: center;
	color: #949494;
	font-size: 55%;
}

.cardInput {
	/* display: inline; */
	width: 20vw;
}

.cardValue-title {
	font-size: 12px;
}

.cardInput {}

/* 动画 */


.guideFst-leave {
	opacity: 1;
}

.guideFst-leave-active {
	transition: all 0.4s;
}

.guideFst-leave-to {
	opacity: 0;
}

.guideSed-enter-active {
	transition: all 0.8s;
}

.guideSed-enter {
	opacity: 0;
}

.guideSed-enter-to {
	opacity: 1;
}
</style>