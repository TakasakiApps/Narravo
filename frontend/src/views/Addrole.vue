<template>
	<div style="margin: 0; padding: 0;">
		<el-carousel height="93vh" direction="horizontal" :autoplay="false" arrow="never" indicator-position="none"
			ref="carousel">
			<el-carousel-item :key="1" class="mian">
				<div v-if="false">
					<h1 class="mian-text">
						暂无角色
						<div>点击右下角按钮以创建
						</div>
					</h1>
				</div>
				<div v-else>
					<RolaCard name="111" :gender=1 Reader="成男1号级" />
				</div>
				<sel-affix position="bottom" :offset="0" class="fixed-element" @click="addRoleCard" v-show="button[0]">
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
							<el-carousel height="65vh" direction="vertical" :autoplay="false" arrow="never" ref="roleCards"
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
										<div class="carValue-main">
											<div class="cardValue-title">语言模型</div>
											<el-select v-model="Card.Reader" class="cardInput" placeholder="选择基本模型" size="default">
												<el-option :label="item[1].name" :value="item[1].name" v-for="item of axgIlem" :key="item[0]"
													@click="readMod(item)" />
											</el-select>
										</div>
										<div class="carValue-main">
											<div class="cardValue-title">默认情感</div>
											<el-select v-model="Card.useemot" class="cardInput" placeholder="选择情绪" size="default">
												<el-option label="不支持" value="no" v-if="!Card.isemotion" />
												<el-option v-for="emotion of Card.emotions" :key="emotion.category" :label="emotion.name"
													:value="emotion.category" v-else />
											</el-select>
										</div>
										<div class="carValue-main">
											<div class="cardValue-title">情感等级</div>
											<el-slider v-model="Card.emotionlv" class="cardInput" :step="1" size="small" show-stops :max="200"
												:disabled="!canEmotion" />
										</div>
										<div class="carValue-main">
											<div class="cardValue-title">语速</div>
											<el-slider v-model="Card.speed" :step="0.5" class="cardInput" size="small" show-stops :max="2.5"
												:min="-2" />
										</div>
										<div class="carValue-main">
											<div class="cardValue-title">音量</div>
											<el-slider v-model="Card.tone" :step="1" class="cardInput" size="small" show-stops :max="10" />
										</div>
									</div>
									<el-button class="btns try-listen" @click="tryListen">试听</el-button>
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
						<div v-show="addTxtIndex" @click="makeRoleCard">
							<svg height="48" viewBox="-30 -750 600 900" width="48" class="btns radiusBtn fsBtn">
								<g transform="scale(0.6)">
									<path d="M378-246 154-470l43-43 181 181 384-384 43 43-427 427Z" fill="currentColor" />
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
import RolaCard from '../components/roleCards/RolaCard.vue';
// import readRoledir from '../components/roleCards/node-api'
import { onMounted, reactive, ref } from 'vue';
import { StarFilled, Plus } from '@element-plus/icons-vue'
import { client as axios } from '../http/client'

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
	//获取全部角色卡
	const token = localStorage.token
	console.log(token)
	if (token) {
		axios.get('/api/v1/tts/models', {
			params: {
				token: token
			}
		})
			.then(
				(e: any) => {
					axgIlem.value = Object.entries(e.data)
					console.log(axgIlem.value)
				}
			)

	}
})
const axgIlem: any = ref([])
const Card = reactive({
	name: '',
	Reader: '',
	sex: 0,
	speed: 0,
	tone: 5,
	isemotion: false,
	useemot: '',
	emotions: [],
	emotionlv: 100,
	modelId: ''
})
const button = reactive([
	true,
	false,
])
const addTxtIndex = ref(false)
//判断是否可以由情绪
const canEmotion = ref('')
const carousel = ref<HTMLFormElement>()
const roleCards = ref<HTMLFormElement>()
function buttonShow() {
	button[0] = !button[0];

}
//根据模型设定值
const readMod = (e: any) => {
	Card.isemotion = e[1].isEmotionSupported;
	Card.sex = e[1].gender;
	Card.emotions = e[1].emotions
	canEmotion.value = e[1].isEmotionSupported
	Card.modelId = e[0]
	console.log(Card)
}
//切换到添加角色
function addRoleCard() {
	if (carousel.value) {
		setTimeout(buttonShow, 500)
		carousel.value.next()
	}

}
//判断当前进度返回到上一级
function backBtn() {
	if (addTxtIndex.value == true) {
		// console.log('返回2-1级')
		if (roleCards.value) {
			addTxtIndex.value = false
			roleCards.value.prev()
		}
	} else {
		// console.log('返回第一级')
		if (carousel.value) {
			buttonShow()
			carousel.value.prev()
		}
	}
}
//返回到上一级
function alterChild() {
	if (Card.name !== '') {
		if (roleCards.value) {
			addTxtIndex.value = true
			roleCards.value.next()
			console.log(addTxtIndex.value)
		}
	} else {
		ElNotification({ message: '名字不能为空！', type: 'warning' })
	}
}
const makeRoleCard = () => {
	if (Card.useemot !== '') {
		console.log('1')
		//重置数据
		Card.name = '';
		Card.Reader = '';
		Card.sex = 0;
		Card.speed = 0;
		Card.tone = 5;
		Card.isemotion = false;
		Card.useemot = '';
		Card.emotions = [];
		Card.emotionlv = 100;
		Card.modelId = ''
	} else {
		ElNotification({ message: '请选择一个模型', type: 'warning' })
	}
}
//试听
const tryListen = () => {

	const data = {
		token: localStorage.token,
		modelId: Card.modelId,
		text: '我超，op！',
		//音量
		volume: Card.tone,
		//语速
		speed: Card.speed,
		codec: 'mp3',
		//使用那种模型
		module: Card.useemot,
		ttsEmotion: {
			enabled: Card.isemotion,
			category: Card.useemot,
			intensity: Card.emotionlv,
		}
	}
	console.log(JSON.stringify(data))
	axios.post('/api/v1/tts/createTask', JSON.stringify(data)).then(
		(e: any) => {
			console.log(e)
		}
	).catch(
		(err: any) => {
			console.log(err)
		}
	)
}
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
	margin-bottom: 4vh;
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
	top: 2.5vh;
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

/* 第二部分设置框 */
.carValue {
	height: calc(52vh - 20px);
}

.cardInput {
	/* display: inline; */
	width: 20vw;
}

.carValue-main {
	position: relative;
	left: 20px;
	top: 20px
}

.cardValue-title {
	font-size: 14px;
	margin-top: 1vh;
	margin-bottom: 1vh;
}

.cardInput {

	width: 19vw;

}

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
	animation: guideSed-in 1s;
	transition: all 0.8s;
}

.guideSed-enter {
	opacity: 0;
}

.guideSed-enter-to {
	opacity: 1;
}

/* 试听按钮样式 */
.try-listen {
	position: relative;
	left: calc(70% - 10px);
	color: white;
	border-radius: 10px;
	font-weight: bold;
}

::v-deep .el-card__body {
	padding: 0px;

}

:deep(.el-input__wrapper) {
	background-color: #f0f0f0;
	border-radius: 10px;
}

/* 取消边缘线 */
::v-deep .el-input__wrapper {
	box-shadow: none !important;
}

::v-deep .el-select .el-input.is-focus .el-input__wrapper {
	box-shadow: none !important;
}

::v-deep .el-select .el-input__wrapper.is-focus {
	box-shadow: none !important;
}

::v-deep .el-select:hover:not(.el-select--disabled) .el-input__wrapper {
	box-shadow: none !important;
}

::v-deep .el-select {
	--el-select-input-focus-border-color: none !important;
}

/* 淡入淡出关键帧 */
@keyframes guideSed-in {
	0% {
		opacity: 0;
	}

	50% {
		opacity: 0;
	}

	100% {
		opacity: 1;
	}
}
</style>