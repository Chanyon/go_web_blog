/**
用vue3重写
by https://github.com/lokyoung/vuejs-paginate/blob/master/src/components/Paginate.vue
*/
<template>
  <div :class="[pageClass, 'paginate']">
    <ul>
      <li v-if="firstLastButton" :class="[pageClass, firstPageSelected() ? disabledClass : '']">
        <a @click="selectFirstPage()" @keyup.enter="selectFirstPage()" :class="pageLinkClass"
          :tabindex="`${firstPageSelected() ? -1 : 0}`" v-html="firstButtonText"></a>
      </li>

      <li v-if="!(firstLastButton && hidePrevNext)" :class="[prevClass, firstPageSelected() ? disabledClass : '']">
        <a @click="prevPage()" @keyup.enter="prevPage()" :class="prevLinkClass"
          :tabindex="`${firstPageSelected() ? -1 : 0}`" v-html="prevText"></a>
      </li>

      <li v-for="page in pages"
        :class="[pageClass, page.selected ? activeClass : '', page.disabled ? disabledClass : '', page.breakView ? breakViewClass : '']">
        <a v-if="page.breakView" :class="[pageLinkClass, breakViewLinkClass]" tabindex="0">
          <slot name="breakViewContent">{{ breakViewText }}</slot>
        </a>
        <a v-else-if="page.disabled" :class="[pageLinkClass]" tabindex="0">{{ page.content }}</a>
        <a v-else @click="handleSetPageSelected(page.index + 1)" @keyup.enter="handleSetPageSelected(page.index + 1)"
          :class="pageLinkClass" tabindex="0">{{ page.content }}</a>
      </li>

      <li v-if="!(firstLastButton && hidePrevNext)" :class="[nextClass, lastPageSelected() ? disabledClass : '']">
        <a @click="nextPage()" @keyup.enter="nextPage()" :class="nextLinkClass"
          :tabindex="`${firstPageSelected() ? -1 : 0}`" v-html="nextText"></a>
      </li>

      <li v-if="firstLastButton" :class="[pageClass, lastPageSelected() ? disabledClass : '']">
        <a @click="selectLastPage()" @keyup.enter="selectLastPage()" :class="pageLinkClass"
          :tabindex="`${firstPageSelected() ? -1 : 0}`" v-html="lastButtonText"></a>
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';
const props = defineProps({
  firstLastButton: {
    type: Boolean,
    default: false,
  },
  hidePrevNext: {
    type: Boolean,
    default: false,
  },
  pageClass: {
    type: String,
    default: "",
  },
  prevClass: String,
  nextClass: String,
  prevLinkClass: String,
  nextLinkClass: String,
  pageLinkClass: String,
  activeClass: {
    type: String,
    default: "active",
  },
  breakViewClass: {
    type: String,
    default: "break-view",
  },
  breakViewLinkClass: {
    type: String,
    default: "break-view-link"
  },
  firstButtonText: {
    type: String,
    default: "first"
  },
  lastButtonText: {
    type: String,
    default: "last"
  },
  prevText: {
    type: String,
    default: "prev",
  },
  nextText: {
    type: String,
    default: "next",
  },
  breakViewText: {
    type: String,
    default: "...",
  },
  clickHandler: {
    type: Function,
    default: () => { },
  },
  pageCount: {
    type: Number,
    default: 8, //8
  },
  pageRange: {
    type: Number,
    default: 3, //3
  },
  marginPages: {
    type: Number,
    default: 1,
  },
});

const selected = ref<number>(1); //current page selected.
const disabledClass = ref<String>("disabled");

// style controller
const firstPageSelected = (): boolean => {
  return selected.value === 1;
}
// style controller
const lastPageSelected = (): boolean => {
  return selected.value === props.pageCount || props.pageCount === 0;
}

const handleSetPageSelected = (selectedPage: number) => {
  // console.log('s' + selected.value);
  // if (selected.value === selectedPage) return;
  selected.value = selectedPage;
  props.clickHandler(selectedPage);
}
const selectFirstPage = () => {
  if (selected.value <= 1) return;
  handleSetPageSelected(1);
};
const selectLastPage = () => {
  if (selected.value > props.pageCount) return;
  handleSetPageSelected(props.pageCount);
};
const prevPage = () => {
  if (selected.value <= 1) return;
  selected.value -= 1;
  handleSetPageSelected(selected.value);
};
const nextPage = () => {
  if (selected.value >= props.pageCount) return;
  selected.value += 1;
  handleSetPageSelected(selected.value);
};

// computed

type PageAndBreakView = {
  [key in number]: {
    index: number,
    content: number,
    selected: boolean,
    disabled: boolean,
    breakView: boolean,
  }
}

const pages = computed(() => {
  let items: PageAndBreakView = {};
  if (props.pageCount <= props.pageRange) {
    for (let idx = 0; idx < props.pageCount; idx++) {
      let page = {
        index: idx,
        content: idx + 1,
        selected: idx === selected.value - 1,
        disabled: false,
        breakView: false,
      };
      items[idx] = page
    }
  } else {
    const halfPageRange = Math.floor(props.pageRange / 2);
    // console.log(halfPageRange);
    const setPageItem = (idx: number) => {
      let page = {
        index: idx,
        content: idx + 1,
        selected: idx === selected.value - 1,
        disabled: false,
        breakView: false,
      };
      items[idx] = page;
    };

    const setBreakView = (idx: number) => {
      let breakView = {
        disabled: true,
        breakView: true,
        index: idx,
        content: idx + 1,
        selected: idx === selected.value - 1,
      };
      items[idx] = breakView;
    };

    //1st loop low end of margin pages
    for (let i = 0; i < props.marginPages; i++) {
      setPageItem(i);
    }

    // 2nd loop thru selected range
    let selectedRangeLow = 0;
    if (selected.value - halfPageRange > 0) {
      selectedRangeLow = selected.value - 1 - halfPageRange;
    }

    let selectedRangeHigh = selectedRangeLow + props.pageRange - 1;
    if (selectedRangeHigh >= props.pageCount) {
      selectedRangeHigh = props.pageCount - 1;
      selectedRangeLow = selectedRangeHigh - props.pageRange + 1;
    }
    for (let i = selectedRangeLow; i <= selectedRangeHigh && i <= props.pageCount; i++) {
      setPageItem(i);
    }

    // check if there is breakView in the left of selected range
    if (selectedRangeLow > props.marginPages) {
      setBreakView(selectedRangeLow - 1);
    }
    if (selectedRangeHigh + 1 < props.pageCount - props.marginPages) {
      setBreakView(selectedRangeHigh + 1);
    }

    // 3rd loop thru high end of margin pages
    for (let i = props.pageCount - 1; i >= props.pageCount - props.marginPages; i--) {
      setPageItem(i);
    }
  }
  return items;
});

</script>

<style scoped>
.paginate a {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  text-align: center;
  /* line-height: 42px; */
  cursor: pointer;
  /* color: white; */
  user-select: none;
}

/* a:hover {
  color: black;
} */

.paginate ul {
  /* margin: 0;
  padding: 0; */
  display: flex;
  justify-content: center;
  align-content: center;
  list-style: none;
  background-color: rgb(255, 116, 106);
  border-radius: 4px;
}

ul>li {
  margin: 4px;
  /* padding: 8px; */
  text-align: center;
  width: 50px;
  height: 42px;
  border: 1px solid #fff;
  border-radius: 8px;
  font-size: 18px;
  cursor: pointer;
}

.disabled {
  background-color: burlywood;
}

.disabled>a {
  cursor: not-allowed;
}

.active {
  background-color: rgb(50, 45, 54);
}

/* props class style */
.page-link {
  color: #fff;
}

.prev-link {
  color: #fff;
}

.next-link {
  color: #fff;
}

.page {
  background-color: rgb(17, 26, 26);
  font-size: 16px;
  border-radius: 8px;
}
</style>