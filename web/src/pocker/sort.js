//  牌排序
export const sortCards = (list) => {
    const getRealVal = (card) => {
        const rank = ["3","4","5","6","7","8","9","10","J","Q","K","A","2","Jack"]
        const pos = rank.indexOf(card.value);
        if(pos === -1)throw new Error("有人出千");
        if(pos === 13 && card.type === 4)return 13;
        if(pos === 13 && card.type === 5)return 14;
        return pos
    }
    for (let i = 0; i < list.length; i++) {
        for (let j = i + 1; j < list.length; j++) {
            if(getRealVal(list[i]) > getRealVal(list[j])){
                [list[j], list[i]] = [list[i], list[j]]
            }
        }
    }
    return list
}

// 出牌排序
export const sortDealCards = (list) => {
    const set = {};
    for(const i in list){
        const card = list[i];
        if(!set[card.value]){
            set[card.value] = [card];
            continue;
        }
        set[card.value].push(card);
    }
    const arr = [];
    for(const i in set){
		sortCards(set[i])
		arr.push(set[i])
	}

    for (let i = 0; i < arr.length; i++) {
        for (let j = i + 1; j < arr.length; j++) {
            if(arr[i].length < arr[j].length){
                [arr[j], arr[i]] = [arr[i], arr[j]]
            }
        }
    }
    const res = []
    for(const i in arr){
        res.push(...arr[i])
    }
    return res
}



