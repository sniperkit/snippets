# Sphinx configuration with Breathe (doxygen)
import os
import sys
import sphinx_rtd_theme
import subprocess

os.chdir(os.path.dirname(__file__) + "/../")

subprocess.call('doxygen', shell=True)

dirname = os.path.abspath(os.path.dirname(__file__) + "/../")
sys.path.insert(0, os.path.join(dirname, "breathe"))
breathe_projects = { "libsyntax7": os.path.join(dirname, "xml")}
breathe_default_project = "libsyntax7"
master_doc = "contents"

extensions = [
	'sphinx.ext.autodoc',
	'sphinx.ext.doctest',
	'sphinx.ext.viewcode',
	'breathe'
]

html_theme = "sphinx_rtd_theme"
html_theme_path = [sphinx_rtd_theme.get_html_theme_path()]

project = 'libsyntax7'
