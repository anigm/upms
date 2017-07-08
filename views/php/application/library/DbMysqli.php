<?php

class DbMysqli
{
    private $error = '';
    private $errno = 0;
    private $port;
    private $host;
    private $username;
    private $password;
    private $dbname;
    private $charset;
    private $db;
    /**
     * 构造函数
     * @author aaron
     * @return void 
     */
    function __construct()
    {
        $config = Yaf_Application::app()->getConfig();
        $this->host = $config->database->params->host; //'127.0.0.1';
        $this->port = $config->database->params->port;  //3306;
        $this->username = $config->database->params->username; //'root';
        $this->password = $config->database->params->password; //'root';
        $this->dbname = $config->database->params->dbname; //'upms';
        $this->charset = $config->database->params->charset; //'UTF8';
        
        $db = mysqli_connect($this->host, $this->username, $this->password, $this->dbname,$this->port);
        if(!$db){
            $this->error = mysqli_connect_error();
            $this->errno = mysqli_connect_errno();
            return FALSE;
        }
        mysqli_set_charset($db, $this->charset);
        
//        $db = mysqli_init();
//        mysqli_options($db, MYSQLI_OPT_CONNECT_TIMEOUT, 5);  // 5s 超时
//        if(!mysqli_real_connect($db, $this->host, $this->username, $this->password, $this->dbname, $this->port)){
//            $this->error = mysqli_connect_error();
//            $this->errno = mysqli_connect_errno();
//            return FALSE;
//        }
//        mysqli_set_charset($db,$this->charset);
        
        $this->db = $db;
    }
    
    
    /**
     * 数据库连接
     * @author aaron
     */
    private function connect()
    {
        return $this->db;
    }
    
    /**
     * 运行Sql语句,不返回结果集
     * @author eric
     * @param string $sql 
     * @return mysqli_result|bool
     */
    public function runSql($sql)
    {
        $dbconnect = $this->connect();
        if ($dbconnect === false) {
            return FALSE;
        }
        // 成功:TRUE 失败:FALSE
        $res = mysqli_query($dbconnect,$sql);
        
        // 保存错误
        $this->save_error($dbconnect);
        
        return $res;
    }
    
    /**
     * 获取受影响的行数
     * @author aaron
     * @param string $sql 
     * @return int
     */
    public function affectedRows()
    {
        $dbconnect = $this->connect();
        if ($dbconnect === false) {
            return FALSE;
        }
        $res = mysqli_affected_rows($dbconnect);
        // 保存错误
        $this->save_error($dbconnect);
        
        return $res;
    }
    
 
    /**
     * 运行Sql,以多维数组方式返回结果集
     * @author aaron
     * @param string $sql 
     * @return array 成功返回数组，失败时返回false
     */
    public function getData($sql)
    {
        // 返回数组
        $data = Array();
        // 初始化连接
        $dbconnect = $this->connect();
        if($dbconnect === false){
            return FALSE;
        }
        // 执行sql语句 返回mysqli_result对象
        $result = mysqli_query($dbconnect,$sql);
        $this->save_error($dbconnect);
        if(is_bool($result)){
            return $result;
        }else{
            while($array = mysqli_fetch_array($result,MYSQLI_ASSOC)){
                $data[] = $array;
            }
        }
        // 关闭 mysqli_result
        mysqli_free_result($result); 
 
        if(count($data)>0){
            return $data;
        }else{
            return NULL;   
        }
    }
    
    
     
    /**
     * 运行Sql,以数组方式返回结果集第一条记录
     * @author aaron
     * @param string $sql 
     * @return array 成功返回数组，失败时返回false
     */
    public function getLine($sql)
    {
        $data = $this->getData($sql);
        if ($data) {
            return @reset($data);
        } else {
            return FALSE;
        }
    }
 
    /**
     * 运行Sql,返回结果集第一条记录的第一个字段值
     * @author aaron
     * @param string $sql 
     * @return mixxed 成功时返回一个值，失败时返回false
     */
    public function getVar( $sql )
    {
        $data = $this->getLine($sql);
        if ($data) {
            return $data[@reset(@array_keys($data))];
        } else {
            return FALSE;
        }
    } 
    
    
    /**
     * 获取新增的id
     * @author aaron
     * @return int 成功返回last_id,失败时返回false
     */
    public function lastId()
    {
        $id = mysqli_insert_id($this->connect());
        return $id;
    }
    /**
     * 关闭数据库连接
     * @author aaron
     * @return bool
     */
    public function closeDb()
    {
        @mysqli_close($this->connect());
    }
    
    
    public function errno()
    {
        return $this->errno;
    }
    public function errmsg()
    {
        return $this->error;
    }
    private function save_error($dbconnect)
    {
        $this->errno = mysqli_errno($dbconnect);
        $this->error = mysqli_error($dbconnect);
    }
}

?>