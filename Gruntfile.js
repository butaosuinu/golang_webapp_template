module.exports = function(grunt) {

	var pkg = grunt.file.readJSON('package.json');

	grunt.initConfig({
		sass:{
			dist:{
				options:{
					style:'expanded'
				},
				files:{
					'public/css/main.css': 'public/sass/main.scss'
				}
			}
		},
		autoprefixer:{
			options:{
				browsers: ['last 2 version', 'ie 9']
			},
			dist:{
				src: ['public/css/main.css']
			}
		},
		cssmin:{
			minify:{
				expand:true,
				src:['public/css/main.css'],
				dest:'',
				ext:'.min.css'
			}
		},
		watch:{
			sass:{
				files:['public/sass/**/*.scss'],
				tasks:['sass','autoprefixer','cssmin'],
			}
		}
	});
	var taskName;
	for (taskName in pkg.devDependencies) {
		if (taskName.substring(0, 6)== 'grunt-') {
			grunt.loadNpmTasks(taskName);
		}
	}
	grunt.registerTask('default', ['watch','sass','autoprefixer','cssmin']);
};
