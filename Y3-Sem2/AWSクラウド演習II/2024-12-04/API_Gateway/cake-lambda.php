<?php
//header('charset=utf-8');
$number = 1001;
//Web APIのURL
$url = 'API GatewayのURL(endpoint)?pid=' . $number;
//APIのJSONファイルを取得
$apidata = file_get_contents($url);
$data = json_decode($apidata, true);
?>
<html lang="ja">
<head>
<!--link rel="stylesheet" href="../css/kad.css" -->
<title>cake shop lambda版</title>
</head>
<body>
<header>
<h1><span>Web API(API Gateway)</span></h1>
</header>
<div id="contents">
<p>
<p>商品</p>
<?php
//表示処理
if(isset($data)){
    print '<p>No.：' . $data['pid'] . '</p>';
    print '<p>商品名：' . $data['pname'] . '</p>';
    print '<p>値段：' . $data['price'] . '円</p>';
}
?>
</div>
</body>
</html>