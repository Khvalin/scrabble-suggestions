var B = 
`⚫⚫⚫A⚫⚫a⚫a⚫⚫A⚫⚫⚫
⚫⚫b⚫⚫B⚫⚫⚫B⚫⚫b⚫⚫
⚫b⚫⚫b⚫⚫⚫⚫⚫b⚫⚫b⚫
A⚫⚫a⚫⚫⚫B⚫⚫⚫a⚫⚫A
⚫⚫b⚫⚫⚫b⚫b⚫⚫⚫b⚫⚫
⚫B⚫⚫⚫a⚫⚫⚫a⚫⚫⚫B⚫
a⚫⚫⚫b⚫⚫⚫⚫⚫b⚫⚫⚫a
⚫⚫⚫B⚫⚫⚫⚫⚫⚫⚫B⚫⚫⚫
a⚫⚫⚫b⚫⚫⚫⚫⚫b⚫⚫⚫a
⚫B⚫⚫⚫a⚫⚫⚫a⚫⚫⚫B⚫
⚫⚫b⚫⚫⚫b⚫b⚫⚫⚫b⚫⚫
A⚫⚫a⚫⚫⚫B⚫⚫⚫a⚫⚫A
⚫b⚫⚫b⚫⚫⚫⚫⚫b⚫⚫b⚫
⚫⚫b⚫⚫B⚫⚫⚫B⚫⚫b⚫⚫
⚫⚫⚫A⚫⚫a⚫a⚫⚫A⚫⚫⚫`;

const lines = B.split('\n')

const res = []
for (const line of lines) {
  const a = []
  for (const ch of line) {
    let item = [1,1]
    if (ch === 'A') {
      item = [1, 3]
    }

    if (ch === 'B') {
      item = [1, 2]
    }
    
    if (ch === 'a') {
      item = [3, 1]
    }
    
    if (ch === 'b') {
      item = [2, 1]
    }
    
    a.push(item)
  }
  
  res.push(a)
}

console.log(JSON.stringify(res) )
