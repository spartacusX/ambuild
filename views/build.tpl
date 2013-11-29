<div class="row">
	<div class="col-lg-22">
		<div class="well">
			<form class="bs-example form-horizontal form-disable" method="post">
				<fieldset>
					<legend>AssetManager Builder</legend>
						<div class="row">
							<label for="platform" class="col-lg-2 control-label">Platform: </label>
							<div class="form-inline">							
								<div class="checkbox col-lg-2">
									<label>
										<input id="platform" name="PlatForm" type="checkbox" value="Windows"> Windows
									</label>
								</div>
								<div class="checkbox col-lg-2">
									<label>
										<input id="platform" name="PlatForm" type="checkbox" value="Linux"> Linux
									</label>
								</div>											
							</div>
						</div>

						<div class="row">
							<label for="encoding" class="col-lg-2 control-label">Encoding: </label>
							<div class="form-inline">							
								<div class="checkbox col-lg-2">
									<label>
										<input id="encoding" name="Encoding" type="checkbox" value="Unicode"> Unicode
									</label>
								</div>
								<div class="checkbox col-lg-2">
									<label>
										<input id="encoding" name="Encoding" type="checkbox" value="ANSI"> ANSI
									</label>
								</div>											
							</div>
						</div>
						
						<div class="row">
							<label for="language" class="col-lg-2 control-label">Language: </label>
							<div class="form-inline">							
								<div class="checkbox col-lg-2">
									<label>
										<input name="Language" type="checkbox" value="All"> All
									</label>
								</div>							
								<div class="checkbox col-lg-2">
									<label>
										<input name="Language" type="checkbox" value="en"> en
									</label>
								</div>
								<div class="checkbox col-lg-2">
									<label>
										<input name="Language" type="checkbox" value="ja"> ja
									</label>
								</div>
								<div class="checkbox col-lg-2">
									<label>
										<input name="Language" type="checkbox" value="zh"> zh
									</label>
								</div>					
							</div>
						</div>

						<div class="row">
							<label for="others" class="col-lg-2 control-label">Advance Options: </label>
							<div class="form-inline">
								<div class="checkbox col-lg-2">
									<label>
										<input name="Sign" type="checkbox"> Sign
									</label>
								</div>
								<div class="checkbox col-lg-2">
									<label>
										<input name="IncBuild" type="checkbox"> Increate Build Number
									</label>
								</div>
								<div class="checkbox col-lg-2">
									<label>
										<input name="MVBuild" type="checkbox"> Move build
									</label>
								</div>							
								<div class="checkbox col-lg-2">
									<label>
										<input name="EnableMOE" type="checkbox"> Enable MOE
									</label>
								</div>
								<div class="checkbox col-lg-2">
									<label>
										<input name="EnableCSA" type="checkbox"> Enable CSA
									</label>
								</div>
							</div>
						</div>
						<input type="submit" class="btn btn-default pull-right {{.BuildStatus}}" value="Launch">
				</fieldset>
			</form>
		</div>
	</div>
</div>
<div class="row col-lg-22">

	<div class="col-lg-6">
		<p><img src="/static/img/building.gif"></p>
	</div>
	<div class="col-lg-6">
		<button id="btn_log" type="button" class="btn btn-default pull-right" onclick="ShowLog()" value="ShowLog">ShowLog</button>
	</div>
</div>

<div id="validation" class="alert alert-warning hidden"></div>

<div class="alert alert-success {{.showhidden}}">
	The build has been launched!
	<br>
	Language: <strong>{{.Language}}</strong>
	Revision: <strong>{{.Revision}}</strong>
	Encoding: <strong>{{.Encoding}}</string>
</div>

<div id="div_log" class="row col-lg-22 hidden">
	<p><a href="http://amdbserver.chn.hp.com/acall/build.php" target="frame_log">Back Log Home</a></p>
	<!-- <button type="button" class="btn btn-info" onclick="RefreshLog()">Log</button> -->
	<!--<iframe id="buildlog" src="/log/prog.html" width="1180" height="900" frameborder="0"></iframe>-->
	<iframe id="buildlog" src="http://amdbserver.chn.hp.com/acall/build.php" name="frame_log" width="1198" height="900" frameborder="0"></iframe>
</div>
<script type="text/javascript">
	function RefreshLog(e){
		var url = "/log";
		//alert(url);
		var xmlhttp;
		if (window.XMLHttpRequest) {// code for IE7+, Firefox, Chrome, Opera, Safari
		  	xmlhttp = new XMLHttpRequest();
		} else {// code for IE6, IE5
		  	xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
		}
		xmlhttp.onreadystatechange = function() 
		{
		  if (xmlhttp.readyState==4 && xmlhttp.status==200) 
		  {
		    try{
		    	document.getElementById('buildlog').innerHTML = xmlhttp.responseText;		    	
		    }catch(er){
		    	document.getElementById('buildlog').innerHTML = "xxxxxxxxx";
		    }
		  }
		}
		xmlhttp.open("GET", url, true);
		xmlhttp.send();
	}

	function RevisionValidate(){
		try{
			var x = document.getElementById("inputRevision").value
			if( isNaN(x) ){
				throw "not a number"
			}
		}catch(err){
			var y = document.getElementById("validation")
			y.className = "alert alert-warning show"
			y.innerHTML = "Error: " + err + "."
			return false
		}

	}

	function ShowLog(){
		try{
			var x = document.getElementById("btn_log")
			var y = document.getElementById("div_log")
			if(x.value == "ShowLog"){
				y.className = "row col-lg-22 show";
				x.value = "HideLog";
				x.innerText = "HideLog";
			}else{
				y.className = "row col-lg-22 hidden";
				x.value = "ShowLog";
				x.innerText = "ShowLog";
			}
			

		}catch(err){
			return false
		}
	}
</script>
