<form class="form-inline" method="post">
		<div class="form-control">
			<div class="form-group col-md-3">
				<label class="control-label" for="selectLanguage">Language:</label>
				<select id="selectLanguage" name="BuildLanguage">
					<option value="en">en</option>
					<option value="zh">zh</option>
					<option value="ja">ja</option>
					<option value="de">de</option>
					<option value="ru">ru</option>
					<option value="ko">ko</option>
					<option value="fr">fr</option>
				</select>
			</div>
			<div class="form-group col-md-3">
				<label class="radio-inline">
				  <input type="radio" name="BuildType" id="optionsRadios1" value="Unicode" checked>
				  Unicode
				</label>
				<label class="radio-inline">
				  <input type="radio" name="BuildType" id="optionsRadios2" value="ANSI">
				  ANSI
				</label>
			</div>
			<div class="form-group col-md-3">
				<label class="control-label" for="inputRevision">Revision: <input type="text" name="BuildRevision" id="inputRevision" placeholder="75803"></label>
			</div>

			<input class="form-group col-md-offset-1" type="submit" value="Launch" />
		</div>
</form>