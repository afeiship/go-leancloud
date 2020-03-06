(function() {
  'use strict';

  const gulp = require('gulp');
  const $ = require('gulp-load-plugins')({
    pattern: ['gulp-*', 'gulp.*', 'del', '@feizheng/gulp-*']
  });

  gulp.task('scripts', function() {
    return gulp
      .src('src/*.js')
      .pipe($.babel())
      .pipe($.uglify())
      .pipe($.feizheng.pkgHeader())
      .pipe(gulp.dest('dist'))
      .pipe($.size({ title: '[ size ]:' }));
  });
})();
