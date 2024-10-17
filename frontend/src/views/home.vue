<template>
  <div class="home-box">
    <div class="home-box-item-l" :class="isShow?'':'no-width'">
      <div class="is-show-btn">
        <a-button size="small" type="primary" shape="circle" @click="isShow=!isShow">
          <MinusOutlined v-if="isShow" />
          <BorderOutlined v-else/>
        </a-button>
      </div>
      <LHeaders v-if="isShow"></LHeaders>
      <LForm v-if="isShow"></LForm>
    </div>
    <div class="home-box-item-r" :class="isShow?'':'full-width',' ',isFull?'full-screen-table':''">
      <RTable @fullScreenTable="fullScreenTable"></RTable>
    </div>
  </div>
</template>

<script>
import { MinusOutlined,BorderOutlined } from '@ant-design/icons-vue'
import RTable from '@/components/rTable.vue'
import LHeaders from '@/components/lHeaders.vue'
import LForm from '@/components/lForm.vue'
export default{
  components:{ RTable,LHeaders,LForm,MinusOutlined,BorderOutlined },
  data(){
    return{
      isShow:true,
      isFull:false
    }
  },
  methods:{
    fullScreenTable(isFull) {
      // console.log(isFull)
      this.isFull=isFull
    }
  }
}
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
  border: 1px solid #f0f0f0;
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
  background: #fff;
  z-index: 999;
  width: 100vw;
}
</style>