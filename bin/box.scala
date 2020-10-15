#! /usr/bin/env scala

import scala.language.postfixOps

val leftUp = "┌"
val leftDown = "└"
val rightUp = "┐"
val rightDown = "┘"
val horizontal = "─"
val vertical = "│"
val empty = " "
val defaultPadding = 0
val defaultConnectors = List.empty[String]

val padding = args.find((e:String) => e.startsWith("-p")).map((e: String) => e.split("-p")(1).toInt).getOrElse(defaultPadding)
val connectors = args.filter((e:String) => e.startsWith("-c")).map((e: String) => e.split("-c")(1))

println("Connectors: " + connectors.mkString(" "))
//TODO refactoring
//TODO connectors > < ^ v
//TODO multiline
//TODO shadow
//TODO empty filler
args.headOption.map { msg: String =>
  val width = 1 + padding + msg.length + padding + 1
  val paddingSpaces = (0 until padding) map (_ => empty) mkString
  val horizontalSide = (0 until width) map (_ => horizontal) mkString
  val emptyLine = (0 until width) map (_ => empty) mkString

  val topLine = leftUp + horizontalSide + rightUp
  val paddingLines = (0 until padding) map { _ =>
    vertical + emptyLine + vertical
  } mkString("\n")

  val contentLine =
    vertical + empty + paddingSpaces + msg + paddingSpaces + empty + vertical

  val bottomLine = leftDown + horizontalSide + rightDown


  val result =
    s"""
      |$topLine
      |${if (paddingLines == "") "" else paddingLines + "\n"}$contentLine${if (paddingLines == "") "" else "\n" + paddingLines}
      |$bottomLine
      |""".stripMargin

  println(result)
}.getOrElse {
  println("Please enter message")
}

