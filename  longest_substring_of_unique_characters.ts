

export default function longestSubstr(s: string): number {
    const letterArray = s.split('');
    let storageArraay = [];
    let current = 0
    let longest = 0

    for(let i = 0; i < letterArray.length; i++){
        for (let j = i; j < letterArray.length; j++){
            if(storageArraay.some(letter => letter === letterArray[j])){
                if(current > longest){
                    longest = current
                }
                current = 0
                storageArraay = []
                break;
            }
            storageArraay.push(letterArray[j])
            current++
        }
    }
    return longest
}




