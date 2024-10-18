<template>
  <a-config-provider :theme="themex" :locale="locale">
    <RouterView />
    <div><About :aboutModal="aboutModal" @closeMoadal="closeMoadal"></About></div>
    <a-float-button-group trigger="hover" type="primary" :style="{ right: '40px',zIndex:999, }">
      <template #icon>
        <SettingOutlined />
      </template>
      <a-float-button href="https://github.com/McaxDev">
        <template #icon>
          <GithubOutlined />
        </template>
      </a-float-button>
      <a-float-button :tooltip="`${themeStore.mode==='dark'?'深':'浅'}色模式`" @click="changeMode(themeStore.mode)">
        <template #icon>
          <icon>
            <svg v-if="themeStore.mode==='light'" xmlns="http://www.w3.org/2000/svg" class="ionicon" viewBox="0 0 512 512"><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="32" d="M256 48v48M256 416v48M403.08 108.92l-33.94 33.94M142.86 369.14l-33.94 33.94M464 256h-48M96 256H48M403.08 403.08l-33.94-33.94M142.86 142.86l-33.94-33.94"/><circle cx="256" cy="256" r="80" fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="32"/></svg>
          </icon>
          <icon>
            <svg v-if="themeStore.mode==='dark'" xmlns="http://www.w3.org/2000/svg" class="ionicon" viewBox="0 0 512 512"><path d="M160 136c0-30.62 4.51-61.61 16-88C99.57 81.27 48 159.32 48 248c0 119.29 96.71 216 216 216 88.68 0 166.73-51.57 200-128-26.39 11.49-57.38 16-88 16-119.29 0-216-96.71-216-216z" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32"/></svg>
          </icon>
        </template>
      </a-float-button>
      <a-float-button :tooltip="'关于'" @click="openModal">
        <template #icon>
          <InfoOutlined />
        </template>
      </a-float-button>
    </a-float-button-group>
  </a-config-provider>
</template>

<script setup>
import { theme } from 'ant-design-vue'
import { ref, watch, onMounted } from 'vue'
import { useThemeStore } from './stores/stores'

const themeStore = useThemeStore()
const { useToken } = theme

const tokens = ref(useToken())

const handleTheme = () => {
  themeStore.handleTheme(tokens.value.token)
}

watch(() => themeStore.mode, (newMode) => {
  handleTheme()
})

onMounted(() => {
  handleTheme()
})
</script>

<script>
import About from './components/about.vue'
import { useThemeStore } from './stores/stores'
import { SettingOutlined, GithubOutlined, InfoOutlined } from '@ant-design/icons-vue'
import zhCN from 'ant-design-vue/es/locale/zh_CN'
import dayjs from 'dayjs'
import 'dayjs/locale/zh-cn'

export default {
  components: {
    SettingOutlined,
    GithubOutlined,
    InfoOutlined,
    About
  },
  data() {
    return {
      themex: {
        token: {
          colorPrimary: '#28abce',
          borderRadius: 8,
        },
        algorithm: theme.defaultAlgorithm,
      },
      locale: zhCN,
      themeStore:useThemeStore(),
      aboutModal:false,
    }
  },
  setup(){
    const themeStore = useThemeStore()
    return {
      themeStore
    }
  },
  mounted(){
    const darkModeMediaQuery = window.matchMedia('(prefers-color-scheme: dark)')
    darkModeMediaQuery.addListener(this.handleDarkModeChange)
    this.handleDarkModeChange(darkModeMediaQuery)
  },
  methods: {
    handleDarkModeChange(e) {
      if (e.matches) {
        this.themeStore.handleMode('dark')
        this.changeMode('dark')
      } else {
        this.themeStore.handleMode('light')
        this.changeMode('light')
      }
    },
    changeMode(mode) {
      if(mode!==''){
        this.themex = {
          token: {
            colorPrimary: '#28abce',
            borderRadius: 8,
          },
          algorithm: this.themeStore.mode === 'dark' ? theme.darkAlgorithm : theme.defaultAlgorithm,
        }
        this.$nextTick(()=>{
          this.themeStore.handleMode(this.themeStore.mode === 'dark' ? 'light' : 'dark')
        })
      }else{
        this.themex = {
          token: {
            colorPrimary: '#28abce',
            borderRadius: 8,
          },
          algorithm: mode === 'light' ? theme.defaultAlgorithm : theme.darkAlgorithm,
        }
      }
    },
    openModal(){
      this.aboutModal=true
    },
    closeMoadal(){
      this.aboutModal=false
    }
  },
}
</script>

<style>
/* 滚动条的宽度和轨道样式 */
::-webkit-scrollbar {
  width:4px;
  height: 4px;
}
::-webkit-scrollbar-track {
  background: #ffffff00;
}

/* 滚动条手柄 */
::-webkit-scrollbar-thumb {
  background: #e6e6e6; 
  border-radius: 4px;
}

::-webkit-scrollbar-thumb:hover {
  background: #e6e6e6;
}
</style>
