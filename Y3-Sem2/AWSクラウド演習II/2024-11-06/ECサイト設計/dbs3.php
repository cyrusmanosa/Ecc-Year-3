<?php
//header('charset=utf-8');
define('HOST', 'sk3a04-db.cjmb5hqdtyry.us-east-1.rds.amazonaws.com');
define('USR', 'main');
define('PASS', '123qwecc');
define('DB', 'sk3a_db');
?>
<html>
<head>
<title>AWS DB test</title>
</head>
<body>
<header><h1>商品一覧</h1></header>
<?php
if (!$conn = mysqli_connect(HOST, USR, PASS, DB)) {
        die('データベースに接続できません');
}

mysqli_set_charset($conn, 'utf8');
$sql = "SELECT * FROM shohins";
$stmt = mysqli_prepare($conn, $sql);
mysqli_stmt_execute($stmt);
mysqli_stmt_bind_result($stmt, $id, $sname, $scategory, $sphoto, $sprice);
print "<table border=\"1\">\n";
print "<tr><td>Photo</td><td>商品名</td><td>カテゴリ</td><td>価格</td></tr>\n";
while (mysqli_stmt_fetch($stmt)) {
    print("<tr>");
    printf("<td><img src=https://sk3a04-bucket.s3.us-east-1.amazonaws.com/images/%s></td>", $sphoto);
    printf("<td>%s</td>", $sname);
    printf("<td>%s</td>", $scategory);
    printf("<td>%d</td>", $sprice);
    print "</tr>\n";
}
print'</table>';
mysqli_stmt_free_result($stmt);
mysqli_stmt_close($stmt);
mysqli_close($conn);
?>
</body>
</html>
