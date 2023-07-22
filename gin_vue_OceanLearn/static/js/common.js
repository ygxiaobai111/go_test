function fetchBackendData1() {
    const form = document.getElementById('login-form');
    const resultContainer = document.getElementById('result-container');
  
    form.addEventListener('submit', function(event) {
      event.preventDefault(); // 阻止表单的默认提交行为
  
      const telephone = document.getElementById('telephone').value;
      const password = document.getElementById('password').value;
  
      // 构造请求体数据
      const formData = new FormData();

      formData.append('telephone', telephone);
      formData.append('password', password);
      // 发送POST请求到登录接口
      fetch('/api/auth/login', {
        method: 'POST',
        body: formData
      })
      .then(response => response.json())
      .then(data => {
        if (data.code === 200) {
          const token = data.data.token;
  
          // 发送GET请求到认证信息接口
          fetch('/api/auth/info', {
            method: 'GET',
            headers: {
              'Authorization': 'Bearer ' + token
            }
          })
          .then(response => response.json())
          .then(data => {
            // 处理返回的数据
            const code = data.code;
            const user_name = data.data.user.name;
            const user_telephone = data.data.user.telephone;
            
            
            // 在result-container中显示返回的信息
            resultContainer.innerHTML = `Code: ${code}, name: ${user_name}, telephone: ${user_telephone}`;
          })
          .catch(error => {
            console.error('Error:', error);
            resultContainer.innerHTML = 'An error occurred while fetching backend data.';
          });
        } else {
          // 处理登录错误
          resultContainer.innerHTML = 'Login failed.';
        }
      })
      .catch(error => {
        console.error('Error:', error);
        resultContainer.innerHTML = 'An error occurred while logging in.';
      });
    });
  }
  
  function fetchBackendData2() {
    const form = document.getElementById('register-form');
    const resultContainer = document.getElementById('result-container2');

    form.addEventListener('submit', function(event) {
        event.preventDefault(); // 阻止默认的表单提交行为
        const name = document.getElementById('name').value;
        const telephone = document.getElementById('telephone').value;
        const password = document.getElementById('password').value;

        const formData = new FormData();
        formData.append('name', name);
        formData.append('telephone', telephone);
        formData.append('password', password);
        

        fetch('/api/auth/register', {
            method: 'POST',
            body: formData
        })
        .then(response => response.json())
        .then(data => {
            // 处理后端返回的数据
            const code = data.code;
            const msg = data.msg;

            // 在页面中显示结果
            resultContainer.innerHTML = `code: ${code}<br>Msg: ${msg}`;
        })
        .catch(error => {
            console.error('Error:', error);
        });
    });
  }
  //字典
function fetchBackendData3(){
 var inputText = document.getElementById("inputText").value;
      var outputContainer = document.getElementById("outputText");
  
      const resultContainer = document.getElementById('result-container');
  
      const formData = new FormData();
      formData.append('inputText', inputText);
  
      fetch("/api/auth/dict", {
        method: "POST",
        body: formData,
      })
      .then(function(response) {
        if (!response.ok) {
          throw new Error("请求失败");
        }
        return response.json();
      })
      .then(function(data) {
        const enUs = data.data.data.dictionary.prons.en_us;
        const en = data.data.data.dictionary.prons.en;
        const explanations = data.data.data.dictionary.explanations;
        const synonym = data.data.data.dictionary.synonym;
        const antonym = data.data.data.dictionary.antonym;
        const wqxExamples = data.data.data.dictionary.wqx_example;
        const entry = data.data.data.dictionary.entry;
        const source = data.data.data.dictionary.source;
  
        const type = data.data.data.dictionary.type;
        const related = data.data.data.dictionary.related;
  
        const formattedOutput = [];
        formattedOutput.push(`<span class="en-us">美: ${enUs}</span>`);
        formattedOutput.push(`<span class="en">英: ${en}</span>`);
        formattedOutput.push(`<span class="explanations">词义: ${explanations}</span>`);
        formattedOutput.push(`<span class="synonym">同义词: ${synonym}</span>`);
        formattedOutput.push(`<span class="antonym">反义: ${antonym}</span>`);
  
        if (wqxExamples.length > 0) {
          formattedOutput.push('<span class="examples">例子:</span>');
          for (let i = 0; i < wqxExamples.length; i++) {
            const exampleStrings = wqxExamples[i].example_strings;
            for (let j = 0; j < exampleStrings.length; j++) {
              formattedOutput.push(`<span class="example">${exampleStrings[j]}</span>`);
            }
          }
        }
  
        formattedOutput.push(`<span class="entry">entry: ${entry}</span>`);
        formattedOutput.push(`<span class="type">type: ${type}</span>`);
        formattedOutput.push(`<span class="related">related: ${related}</span>`);
        formattedOutput.push(`<span class="source">source: ${source}</span>`);
  
        outputContainer.innerHTML = formattedOutput.join("<br>");
      })
      .catch(function(error) {
        console.log(error);
      });
}

  

