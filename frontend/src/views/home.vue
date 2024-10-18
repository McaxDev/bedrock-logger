<template>
  <div class="home-box" :style="{ background: theme.main.colorBgLayout }">
    <div v-if="!isFull" class="home-box-item-l" :class="isShow?'':'no-width'" :style="{ borderColor: theme.main.colorBorder }">
      <div class="is-show-btn">
        <a-button size="small" type="primary" shape="circle" @click="isShow=!isShow">
          <MinusOutlined v-if="isShow" />
          <BorderOutlined v-else/>
        </a-button>
      </div>
      <LHeaders v-show="isShow"></LHeaders>
      <LForm v-show="isShow"></LForm>
    </div>
    <div class="home-box-item-r" :class="isShow?'':'full-width',' ',isFull?'full-screen-table':''">
      <RTable @fullScreenTable="fullScreenTable"></RTable>
    </div>
  </div>
</template>

<script>
import { defineComponent, computed, watch, ref } from 'vue'
import { MinusOutlined,BorderOutlined } from '@ant-design/icons-vue'
import RTable from '@/components/rTable.vue'
import LHeaders from '@/components/lHeaders.vue'
import LForm from '@/components/lForm.vue'

import { useThemeStore } from '@/stores/stores'

export default defineComponent({
  components:{ RTable,LHeaders,LForm,MinusOutlined,BorderOutlined },
  data(){
    return{
      isShow:true,
      isFull:false,
    }
  },
  setup() {
    const theme = useThemeStore()
    return {
      theme
    }
  },
  mounted(){
    
  },
  methods:{
    fullScreenTable(isFull) {
      // console.log(isFull)
      this.isFull=isFull
    }
  },
  watch:{

  }
})
</script>

<style scoped>
.home-box{
  display: flex;
  justify-content: space-between;
  padding: 20px;
  min-height: 100vh;
}
.home-box-item-l{
  position: relative;
  width: 40%;
  border-radius: 8px;
  padding: 20px 10px;
  border: 1px solid;
  max-height: calc(100vh - 40px);
  overflow: auto;
  transition: 0.2s;
}
.home-box-item-r{
  width: 60%;
  transition: 0.2s;
}
.is-show-btn{
  position: absolute;
  top: 10px;
  right: 10px;
}
.no-width{
  padding: 0;
  margin: 0;
  width: 44px;
  height: 44px;
  border:none;
  transition: 0.2s;
}
.full-width{
  width: calc(100% - 44px);
}
.full-screen-table{
  transition: 0.2s;
  position: fixed;
  top: 0;
  left: 0;
  padding: 10px;
  min-height: 100vh;
  overflow: auto;
  z-index: 999;
  width: 100vw;
}
@media (max-width: 992px) {
  .home-box{
    display: flex;
    flex-direction: column;
    padding: 20px;
    min-height: 100vh;
    max-height: none;
    justify-content: flex-start;
  }
  .home-box-item-l{
    position: relative;
    width: 100%;
    max-height: none;
    overflow: auto;
    transition: 0.2s;
    margin-bottom: 10px;
  }
  .home-box-item-r{
    width: 100%;
    transition: 0.2s;
  }
  .full-screen-table{
    min-height: none;
  }
}
</style>

<style>
@media (max-width: 992px) {
  .ant-picker-datetime-panel{
    display: flex;
    flex-direction: column;
  }
}
</style>